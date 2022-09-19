/*
 * @Author: Jerry.Yang
 * @Date: 2022-09-07 14:34:01
 * @LastEditors: yangjie04@qutoutiao.net
 * @LastEditTime: 2022-09-19 19:06:24
 * @Description: create app
 */
package operate

import (
	"errors"
	"fmt"
	"strings"

	mytoolpkg "github.com/yangjerry110/mytool/pkg"
)

type (
	CreateAppInterface interface{}

	CreatedApp struct{}
)

/**
 * @description: Action
 * @param {string} operateParam
 * @author: Jerry.Yang
 * @date: 2022-09-07 14:37:39
 * @return {*}
 */
func (c *CreatedApp) Action(operateParam string) {

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
		fmt.Printf("createApp Err : %v", err.Error())
		fmt.Print("\r\n")
		return
	}

	/**
	 * @step
	 * @定义参数
	 **/
	projectName := appParams[0] // project
	appName := appParams[1]     // appName
	method := appParams[2]      // method

	/**
	 * @step
	 * @创建app
	 **/
	mytoolpkg.CreateApp(projectName, appName, method)
	return
}

/**
 * @description: CheckAppParams
 * @param {[]string} appParams
 * @author: Jerry.Yang
 * @date: 2022-09-07 14:47:04
 * @return {*}
 */
func (c *CreatedApp) CheckAppParams(appParams []string) error {

	/**
	 * @step
	 * @判断参数
	 **/
	if len(appParams) == 0 {
		return errors.New("projectName缺失!")
	}

	if len(appParams) == 1 {
		return errors.New("appName确实!")
	}

	if len(appParams) == 2 {
		return errors.New("请求方式缺失!")
	}
	return nil
}
