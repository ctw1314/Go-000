package main

import (
	"errors"
	"fmt"
	xerrors "github.com/pkg/errors"
)

/*
作业：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
答案：
一般应该wrap这个error，因为这是第三方库的报错，dap层wrap，业务层直接返回，顶层输出，这样错误信息完整、清晰、不冗余
如果有某一些特定场景下认为不应该报错，那么应该在该业务的service层内单独消化这种情况
如下所示
*/

//main
func main() {
	err := service()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}

//dao
func dao() error {
	// 模拟dao遇到sql.ErrNoRows
	ErrorNoRows := errors.New("sql.ErrorNoRows")
	return xerrors.Wrap(ErrorNoRows, "dao error")
}

//biz
func biz() error {
	return dao()
}


//service
func service() error {
	return biz()
}


