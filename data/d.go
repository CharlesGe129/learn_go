package data

var datas []string

func Add(str string) string {
	datas = append(datas, str)

	return str
}