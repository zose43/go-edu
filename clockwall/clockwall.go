package clockwall

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

func ClockWall(servers map[string]int) {
	clockWall := make(map[string]string)
	for tz, port := range servers {
		t, err := handle(port)
		if err != nil {
			log.Print(fmt.Sprintf("%s %d %v", tz, port, err))
			continue
		}
		clockWall[tz] = t
	}

	if len(clockWall) > 0 {
		const format = "%s\t%s\n"
		tw := new(tabwriter.Writer).Init(os.Stdout, 0, 0, 4, ' ', 0)
		_, _ = fmt.Fprintf(tw, format, "Timezone", "Time")
		_, _ = fmt.Fprintf(tw, format, "------", "------")
		for zone, time := range clockWall {
			_, _ = fmt.Fprintf(tw, format, zone, time)
		}
		err := tw.Flush()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func handle(port int) (string, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return "", err
	}
	defer func() { _ = conn.Close() }()

	r := bufio.NewReader(conn)
	tz, err := r.ReadString('\n')

	if err != nil {
		return "", fmt.Errorf("can't read body %w", err)
	}
	return strings.TrimSpace(tz), nil
}

//servers := map[string]int{
//"US/Eastern":    8000,
//"Asia/Tokyo":    8001,
//"Europe/London": 8002,
//}
