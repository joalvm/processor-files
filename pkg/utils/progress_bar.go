package utils

import "github.com/cheggaaa/pb/v3"

func ProgressBar(size int) *pb.ProgressBar {
	bar := pb.New(size)

	bar.SetTemplateString(`{{counters . }} {{bar . "|" (green "█") (white "▓") (red "░") "|"}} {{percent . }} {{string . "filename" }}`)

	return bar
}
