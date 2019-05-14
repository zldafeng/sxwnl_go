package src

import (
	"fmt"
	"testing"
)

/**
 * @Author: Gale 598787663@qq.com
 * @Date: 19:40 2019-05-14
 * @Description:
 */
func TestGetDayBySolar(t *testing.T) {
	d := GetDayBySolar(2019,5,14)
	fmt.Println(d)
}