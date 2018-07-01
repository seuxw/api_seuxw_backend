package extension

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//获取随机数Code
func GetRandCode(prefix string) string {
	randcode := prefix + strconv.FormatInt(time.Now().Unix(), 10)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%04v", rnd.Int31n(10000))
	return randcode + vcode
}
