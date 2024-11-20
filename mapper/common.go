package mapper

import (
	"context"
	"reflect"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type ctxTransactionKey struct{}

func SetDB(tx *gorm.DB) {
	db = tx
	_ = registerCallback(db)
}

func GetDB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})
	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if !ok {
			zap.L().Panic("unexpect context value type", zap.String("type", reflect.TypeOf(tx).String()))
			return nil
		}
		return tx
	}
	return db.WithContext(ctx)
}

func Tx(ctx context.Context, fn func(txctx context.Context) error) error {
	db := db.WithContext(ctx)
	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fn(txctx)
	})
}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxTransactionKey{}, tx)
}

func registerCallback(tx *gorm.DB) error {

	if err := tx.Callback().Create().Before("gorm:create").Register("created", created); err != nil {
		return err
	}

	if err := tx.Callback().Update().Before("gorm:update").Register("updated", updated); err != nil {
		return err
	}

	return nil
}

func created(tx *gorm.DB) {

	nowTime := time.Now()
	if field, ok := tx.Statement.Schema.FieldsByName["CreatedAt"]; ok {
		kind(tx, field, func() {
			tx.Statement.SetColumn("CreatedAt", nowTime, true)
		})
	}

	if field, ok := tx.Statement.Schema.FieldsByName["UpdatedAt"]; ok {
		kind(tx, field, func() {
			tx.Statement.SetColumn("UpdatedAt", nowTime, true)
		})
	}

	// 获取当前登录用户
	// user, err := auth.GetUserCtx(tx.Statement.Context)
	// if err != nil {
	// 	return
	// }

	// if field, ok := tx.Statement.Schema.FieldsByName["CreatedBy"]; ok {
	// 	kind(tx, field, func() {
	// 		tx.Statement.SetColumn("CreatedBy", user.Username, true)
	// 	})
	// }

}

func updated(tx *gorm.DB) {

	if field, ok := tx.Statement.Schema.FieldsByName["UpdatedAt"]; ok {
		kind(tx, field, func() {
			tx.Statement.SetColumn("UpdatedAt", time.Now(), true)
		})
	}

	// 获取当前登录用户
	// user, err := auth.GetUserCtx(tx.Statement.Context)
	// if err != nil {
	// 	return
	// }

	// if field, ok := tx.Statement.Schema.FieldsByName["UpdatedBy"]; ok {
	// 	kind(tx, field, func() {
	// 		tx.Statement.SetColumn("UpdatedBy", user.Username, true)
	// 	})
	// }

}

func kind(tx *gorm.DB, field *schema.Field, f func()) {
	switch tx.Statement.ReflectValue.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < tx.Statement.ReflectValue.Len(); i++ {
			// Get value from field
			if _, isZero := field.ValueOf(tx.Statement.Context, tx.Statement.ReflectValue.Index(i)); isZero {
				f()
			}
		}
	case reflect.Struct:
		// Get value from field
		if _, isZero := field.ValueOf(tx.Statement.Context, tx.Statement.ReflectValue); isZero {
			f()
		}
	}
}
