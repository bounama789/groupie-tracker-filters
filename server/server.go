package server

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/helper"
	"groupie-tracker/models"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var ErrorMsgMap = map[int]Error{
	http.StatusBadRequest:          {http.StatusBadRequest, "Bad Request"},
	http.StatusNotFound:            {http.StatusNotFound, "Not Found"},
	http.StatusInternalServerError: {http.StatusInternalServerError, "Internal Server Error"},
	http.StatusMethodNotAllowed:    {http.StatusMethodNotAllowed, "Method Not Allowed"},
}

type Error struct {
	Code int
	Msg  string
}

func getAllArtists() ([]models.Artist, error) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var artists []models.Artist

	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func getAllLocations() (models.AllLocations, error) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	var loca models.AllLocations

	if err != nil {
		return loca, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&loca)
	if err != nil {
		return loca, err
	}

	return loca, nil
}

func getAllConcertDates() ([]models.Dates, error) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var dates []models.Dates

	err = json.NewDecoder(response.Body).Decode(&dates)
	if err != nil {
		return nil, err
	}

	return dates, nil
}

func getArtistById(id int, artist *models.Artist) error {

	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(artist)
	if artist.Id == 0 {
		err = fmt.Errorf("no artist with this ID: %d", id)
	}

	if err != nil {
		return err
	}

	return err
}

func getOneLocations(url string, location *models.Location) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(location)
	if err != nil {
		return err
	}

	return err
}

func getOneDates(url string, date *models.Dates) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(date)
	if err != nil {
		return err
	}

	return err
}

func getOneRelations(url string, relations *models.Relations) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	err = json.Unmarshal(body, &relations)
	if err != nil {
		return err
	}

	return err
}

type Res struct {
	Artists []models.Artist
	Locations models.AllLocations
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	locations ,_ := getAllLocations()
	if r.URL.Path != "/" {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	if r.Method != http.MethodGet {
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
	artists, err := getAllArtists()
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	s := "templates/index.html"
	tmpl, err := template.ParseFiles(s)

	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}

	res := Res{
		artists,
		locations,
	}


	tmpl.Execute(w, res)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}
	if r.Method != http.MethodGet {
		RenderErrorPage(http.StatusMethodNotAllowed, w)
		return
	}
	s := "templates/artist.html"
	tmpl, err := template.ParseFiles(s)
	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id < 0 {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}
	var artist models.Artist

	err = getArtistById(id, &artist)

	if err != nil {
		RenderErrorPage(http.StatusNotFound, w)
		return
	}

	var artistRes = generateArtistResponse(artist)

	tmpl.Execute(w, artistRes)

}

// hereeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee

// type FilterData struct {
// 	Id           int
// 	Image        string
// 	Name         string
// 	Members      []string
// 	CreationDate int
// 	FirstAlbum   string
// 	Locations    string
// 	Relations    string
// }

func FilterByCreationDate(min int, max int ,arrayFilter []models.Artist) []models.Artist {
	//artists,_ := getAllArtists()
	 var filterData []models.Artist
	 println(min,max)
	for _, item := range arrayFilter {
		if (item.CreationDate >= min) && (item.CreationDate <= max)  {
			println(item.CreationDate)
			filterData = append(filterData, item)
		}
	}

	return filterData
}
func FilterByFirstAlbum(min int, max int, arrayFilter []models.Artist) []models.Artist {
	//artists,_ := getAllArtists()
	 var filterData []models.Artist
	 println(min,max)
	for _, item := range arrayFilter {
		tab := strings.Split((item.FirstAlbum),"-")
		date,_ := strconv.Atoi(tab[len(tab)-1])
		if (date >= min) && (date <= max)  {
			println(item.CreationDate)
			filterData = append(filterData, item)
		}
	}

	return filterData
}
func FilterByMember(number []int, arrayFilter []models.Artist ) []models.Artist  {
	
	//artists,_ := getAllArtists()
	var filterData []models.Artist

	if len(number) > 0 {
		
		
			for _, item := range arrayFilter {
				for _, v := range number {
					if len(item.Members) == v {
					filterData = append(filterData, item)
				}
				}
				
			}
			return filterData
	}
return arrayFilter

}
func FilterByLocation(str string, arrayFilter []models.Artist) []models.Artist {
	locations ,_ := getAllLocations()

	res := Res{
		arrayFilter,
		locations,
	}

	var filterData []models.Artist

	if str=="all" {
		return arrayFilter
		
	}

	for i, item := range res.Locations.Index {
		for _, v := range item.Locations {
			//fmt.Println(v)
			//fmt.Println(str)
			if v==str {
				fmt.Printf(v,str)
				filterData = append(filterData, arrayFilter[i])
			}
			
		}
		
		
	}

	return filterData
}
func HandlerFilter(w http.ResponseWriter, r *http.Request){
	location ,_ := getAllLocations()
	artists,_ := getAllArtists()
	s := "templates/index.html"
	tmpl, err := template.ParseFiles(s)

	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
	min,_ := strconv.Atoi(r.FormValue("min"))
	max,_ := strconv.Atoi(r.FormValue("max"))
	minbis,_ := strconv.Atoi(r.FormValue("minbis"))
	maxbis,_ := strconv.Atoi(r.FormValue("maxbis"))
	option := r.FormValue("option")
	checkboxValue1,_ := strconv.Atoi(r.FormValue("checkboxValue1"))
	checkboxValue2,_ := strconv.Atoi(r.FormValue("checkboxValue2"))
	checkboxValue3,_ := strconv.Atoi(r.FormValue("checkboxValue3"))
	checkboxValue4,_ := strconv.Atoi(r.FormValue("checkboxValue4"))
	checkboxValue5,_ := strconv.Atoi(r.FormValue("checkboxValue5"))
	checkboxValue6,_ := strconv.Atoi(r.FormValue("checkboxValue6"))
	checkboxValue7,_ := strconv.Atoi(r.FormValue("checkboxValue7"))
	checkboxValue8,_ := strconv.Atoi(r.FormValue("checkboxValue8"))

	fmt.Println(option)
	
	var numbers = []int{
		checkboxValue1,
		checkboxValue2,
		checkboxValue3,
		checkboxValue4,
		checkboxValue5,
		checkboxValue6,
		checkboxValue7,
		checkboxValue8,
	}
	var tabNumbers = []int{}

	for _, v := range numbers {
		if v!=0 {
			tabNumbers = append(tabNumbers, v)
			fmt.Println(v)
		}
	}
	
	filterRes := FilterByCreationDate(min, max, artists)
	filterRes = FilterByFirstAlbum(minbis, maxbis, filterRes)
	filterRes = FilterByLocation(option, filterRes)
	filterRes = FilterByMember(tabNumbers, filterRes)
	 res := Res{
		filterRes,
		location,
	}
	//members := FilterByMember()


	tmpl.Execute(w, res)
}
}

func SuggestHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	artists, err := getAllArtists()

	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}
	locations, err := getAllLocations()

	searchReq := strings.ToLower(r.Form.Get("q"))

	if err != nil {
		RenderErrorPage(http.StatusInternalServerError, w)
		return
	}

	var suggest models.SuggestResp

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), searchReq) {
			suggest.Bands = helper.AppendIfNotExist(suggest.Bands, artist.Name)
		}

		if strings.Contains(strconv.Itoa(artist.CreationDate), searchReq) {
			suggest.CreationDates = helper.AppendIfNotExist(suggest.CreationDates, strconv.Itoa(artist.CreationDate))
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), searchReq) {
				suggest.Members = helper.AppendIfNotExist(suggest.Members, member)
			}

		}

		for _, location := range locations.Index {
			helper.FormatLocations(&location)
			for _, loc := range location.Locations {
				if strings.Contains(strings.ToLower(loc), strings.ToLower(searchReq)) {
					suggest.Locations = helper.AppendIfNotExist(suggest.Locations, loc)
				}
			}
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), searchReq) {
			suggest.FirstAlbums = helper.AppendIfNotExist(suggest.FirstAlbums, artist.FirstAlbum)
		}

	}

	htmlElem := `
	{{range $i, $name := .Bands}}
	<option value="{{$name}}">{{$name}} - artist/band </option>
	{{end}}

	{{range $i, $member := .Members}}
	<option value="{{$member}}">{{$member}} - member </option>
	{{end}}

	{{range $i, $firstAlbum := .FirstAlbums}}
	<option value="{{$firstAlbum}}">{{$firstAlbum}} - first_album</option>
	{{end}}

	{{range $i, $location := .Locations}}
	<option value="{{$location}}">{{$location}} - location </option>
	{{end}}

	{{range $i, $creationDate:= .CreationDates}}
	<option value="{{$creationDate}}">{{$creationDate}} - creation_date </option>
	{{end}}`

	tmpl, _ := template.New("sugg").Parse(htmlElem)

	if strings.TrimSpace(searchReq) != "" && len(searchReq) > 0 {
		tmpl.Execute(w, suggest)
	}

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		RenderErrorPage(http.StatusBadRequest, w)
		return
	}

	qSplit := strings.Split(q, " - ")

	if len(qSplit) == 2 {
		q = qSplit[0]
	}

	artists, err := getAllArtists()
	if err != nil {
		RenderErrorPage(http.StatusBadGateway, w)
		return
	}

	var artistsResp []models.Artist
	locations, err := getAllLocations()

	if err != nil {
		RenderErrorPage(http.StatusBadGateway, w)
		return
	}

	for _, artist := range artists {
		found := false
		if strings.HasPrefix(strings.ToLower(artist.Name), q) {
			artistsResp = append(artistsResp, artist)
			continue
		}

		if strings.HasPrefix(strconv.Itoa(artist.CreationDate), q) {
			artistsResp = append(artistsResp, artist)
			continue

		}

		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), q) {
				artistsResp = append(artistsResp, artist)
				found = true
				break
			}

		}
		if found {
			continue
		}

		if found {
			continue
		}

		if strings.HasPrefix(strings.ToLower(artist.FirstAlbum), q) {
			artistsResp = append(artistsResp, artist)
			continue
		}
	}

	for _, location := range locations.Index {
		helper.FormatLocations(&location)
		for _, loc := range location.Locations {
			if helper.HasKeyword(strings.ToLower(q), strings.ToLower(loc)) {
				var artist models.Artist
				getArtistById(location.Id, &artist)
				artistsResp = append(artistsResp, artist)
				break
			}
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, artistsResp)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/static/")

	if path != "" {
		http.ServeFile(w, r, "./static/"+path)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func RenderErrorPage(errorCode int, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/error.html")

	if err != nil {
		w.WriteHeader(errorCode)
		tmpl.Execute(w, ErrorMsgMap[http.StatusInternalServerError])
		return
	}

	w.WriteHeader(errorCode)
	tmpl.Execute(w, ErrorMsgMap[errorCode])
}

func generateArtistResponse(artist models.Artist) models.ArtistResponse {
	var artistRes models.ArtistResponse

	artistRes.Id = artist.Id
	artistRes.CreationDate = artist.CreationDate
	artistRes.FirstAlbum = helper.FormatDate(artist.FirstAlbum)
	artistRes.ImageURL = artist.Image
	artistRes.Members = artist.Members
	artistRes.Name = artist.Name

	getOneLocations(artist.Locations, &artistRes.Locations)
	getOneDates(artistRes.Locations.Dates, &artistRes.ConcertDates)

	artistRes.TotalConcerts = len(artistRes.ConcertDates.Dates)
	getOneRelations(artist.Relations, &artistRes.Relations)
	helper.FormatConcertDates(&artistRes.Relations)

	return artistRes
}
