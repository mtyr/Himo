package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mtyr/Himo/matuura/model"
	"github.com/mtyr/Himo/matuura/repo"
	"github.com/tarm/serial"
)

var ms byte
var flag int

type Test struct {
	Title string    `json:"title"`
	Time  time.Time `json:"time"`
	Slice [][]int   `json:"slice"`
}

// time 2018-08-16T17:16:08.454+09:00

func main() {
	go comm()
	//引数: ファイルのパス, フラグ, パーミッション(わからなければ0666でおっけーです)
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintln(file, "あかしけ　やなげ　緋色の鳥よ　くさはみ　ねはみ　けをのばせ ") //書き込み
	/*
		r := gin.Default()
		c := &serial.Config{Name: "COM8", Baud: 9600}
		s, err := serial.OpenPort(c)

		if err != nil {
			// jsonデータを出力
			//r.GET("/Situation", func(c *gin.Context) {
			//	c.JSON(http.StatusOK, Situ)
			//})
			//log.Fatal(err)
			flag = 2
		}

		r.GET("/Situation", func(c *gin.Context) {
			c.JSON(http.StatusOK, flag)
		})
		if flag == 2 {
			r.Run()
		}
		n, err := s.Write([]byte("@"))
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, 32)
		//for i := 0; i < 32; i++ {
		//n, err = s.Read(buf)
		//}
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
			//fmt.println(n)
		}
	*/
	//fmt.Printf("%q", buf[:n])

	r := gin.Default()
	fmt.Println(time.Now())
	r.GET("/test", func(c *gin.Context) {
		var test Test
		test.Title = "matuura"
		test.Time = time.Now()
		test.Slice = [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
		//test.Slice = [][]int{{}, {}}
		//var b[] int = int(buf[])
		//Bytes2int(buf[])
		//test.Slice = [][]byte{buf}
		c.JSON(http.StatusOK, test)
	})

	r.POST("/alarm", func(c *gin.Context) {
		var alarm model.Alarm
		err := c.BindJSON(&alarm)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		repo.SetAlarm(alarm)
	})

	r.GET("/alarms", func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.GetAllAlarm())
	})

	r.GET("/alarm/:seet_id", func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.FindBySeetIDAlarm(c.Param("seet_id")))
	})
	go r.Run()

	for {
	}

}

func comm() {
	r := gin.Default()
	c := &serial.Config{Name: "COM8", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		//log.Fatal(err)
		flag = 2
	}

	r.GET("/Situation", func(c *gin.Context) {
		c.JSON(http.StatusOK, flag)
	})
	if flag == 2 {
		r.Run()
	}

	r.GET("/alarms", func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.GetAllAlarm())
	})

	r.GET("/Seat", func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.SetSeat())
	})

	n, err := s.Write([]byte("test"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 32)
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 32; i++ {
		fmt.Printf("%q", buf[i])
	}
	fmt.Printf("\n")
	fmt.Printf("%q\n", buf[:n])

}

/*
	ms = 0
	c := &serial.Config{Name: "COM6", Baud: 9600}
	s, err := serial.OpenPort(c)
	r := gin.Default()
	if err != nil {
		//log.Fatal(err)
		flag = 2
	}
	n, err := s.Write([]byte("@"))
	if err != nil {
		log.Fatal(err)
	}
	r.GET("/Situation", func(c *gin.Context) {
		c.JSON(http.StatusOK, flag)
	})
	if flag == 2 {
		r.Run()
	}
	buf := make([]byte, 32)
	seat := make([]int, 0)
	for {
		n, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		print("%d", buf)
		for i := 0; i < 32; i++ {
			if buf[i] == ms {
				seat = append(seat, 0)
			} else {
				seat = append(seat, 1)
			}
		}
	}
}*/

//fmt.Printf("%q", buf[:n])
//}
//n, err = s.Read(buf)
//if err != nil {
//	log.Fatal(err)
//}
//fmt.Printf("%q", buf[:n])
//}
/*
func Bytes2int(bytes ...byte) int64 {
	if 0x7f < bytes[0] {
		mask := uint64(1<<uint(len(bytes)*8-1) - 1)

		bytes[0] &= 0x7f
		i := Bytes2uint(bytes...)
		i = (^i + 1) & mask
		return int64(-i)

	} else {
		i := Bytes2uint(bytes...)
		return int64(i)
	}
}*/
