package objloader

import (
	"os"
	"bufio"
	"fmt"
)

type MeshObject struct {
	Vertices []float32 
	UVs []float32
	Normals []float32
}

func LoadObject(fname string) *MeshObject {
	file,err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	vertices, uvs, normals :=  make([][3]float32,0),make([][2]float32,0),make([][3]float32,0)
	vIndices, uvIndices, nIndices := make([]uint,0), make([]uint,0), make([]uint,0)

	for line,err := reader.ReadString('\n'); err == nil; line,err = reader.ReadString('\n') {
		lineReader := bufio.NewReader(StringReader(line))
		header,_ := lineReader.ReadString(' ')
		restOfLine,_ := lineReader.ReadString('\n')

		switch header[:len(header)-1] {
		case "v": 
			vert := [3]float32{}
			count,_ := fmt.Sscanf(restOfLine, "%f %f %f\n", &vert[0], &vert[1], &vert[2])
			if count != 3 {
				panic("Wrong vert count")
			}
			vertices = append(vertices, vert)
		
		case "vt":
			uv := [2]float32{}
			count,_ := fmt.Sscanf(restOfLine, "%f %f\n", &uv[0], &uv[1] )
			if count != 2 {
				panic("Wrong uv count")
			}
			uvs = append(uvs, uv)
		case "vn":
			norm := [3]float32{}
			count,_ := fmt.Sscanf(restOfLine, "%f %f %f\n", &norm[0], &norm[1], &norm[2])
			if count != 3 {
				panic("Wrong norm count")
			}
			normals = append(normals, norm)
		case "f":
			//vert1, vert2, vert3 string
			vIndex, uvIndex, nIndex := [3]uint{}, [3]uint{}, [3]uint{}
			matches,_ := fmt.Sscanf(restOfLine, "%d/%d/%d %d/%d/%d %d/%d/%d\n", &vIndex[0], &uvIndex[0], &nIndex[0], &vIndex[1], &uvIndex[1], &nIndex[1], &vIndex[2], &uvIndex[2], &nIndex[2])
			if matches != 9 {
				panic("Can't read file")
			}
			vIndices = append(vIndices, vIndex[:]...)
			uvIndices = append(uvIndices, uvIndex[:]...)
			nIndices = append(nIndices, nIndex[:]...)
		default:
			// eat line
		}
	}

	//fmt.Println(vertices)
	
	obj := &MeshObject{make([]float32,0,len(vIndices)*3),make([]float32,0,len(uvIndices)*2),make([]float32,0,len(nIndices)*3) }
	for i := range vIndices {
		vIndex, uvIndex, nIndex := vIndices[i], uvIndices[i], nIndices[i]

		vert, uv, norm := vertices[vIndex-1], uvs[uvIndex-1], normals[nIndex-1]

		obj.Vertices = append(obj.Vertices, vert[:]...)
		obj.UVs = append(obj.UVs, uv[:]...)
		obj.Normals = append(obj.Normals, norm[:]...)
	}
	
	return obj
}

type StringReader string

func (s StringReader) Read(byt [] byte) (n int, err error) {
	copy(byt,string(s))
	return len(byt), nil
}
