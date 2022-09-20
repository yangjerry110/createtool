/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-13 18:47:16
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-20 10:43:30
 * @Description: create dao
 */
package operate

import (
	"errors"
	"fmt"
	"strings"

	mytoolpkg "github.com/yangjerry110/mytool/pkg"
)

type (
	CreateDao struct{}

	CreateDaoInfo struct{}
)

func (c *CreateDao) Action(operateParam string) {
	/**
	 * @step
	 * @解析参数
	 **/
	appParams := strings.Split(operateParam, ".")

	/**
	 * @step
	 * @检查参数
	 **/
	err := c.CheckAppParams(appParams)
	if err != nil {
		fmt.Printf("CreateDao Err : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @获取authorName
	 **/
	if len(appParams) == 3 {
		appParams = append(appParams, "Jerry.Yang")
	}

	/**
	 * @step
	 * @定义参数
	 **/
	projectName := appParams[0]
	daoName := appParams[1]
	modelName := appParams[2]
	authorName := appParams[3]

	/**
	 * @step
	 * @创建dao
	 **/
	mytoolpkg.CreateDao(projectName, daoName, modelName, authorName)
	return
}

/**
 * @description: CheckAppParams
 * @param {[]string} appParams
 * @author: Jerry.Yang
 * @date: 2022-09-14 16:11:57
 * @return {*}
 */
func (c *CreateDao) CheckAppParams(appParams []string) error {

	/**
	 * @step
	 * @判断参数
	 **/
	if len(appParams) == 0 {
		return errors.New("projectName缺失!")
	}

	if len(appParams) == 1 {
		return errors.New("daoName缺失!")
	}

	if len(appParams) == 2 {
		return errors.New("modelName缺失!")
	}
	return nil
}
