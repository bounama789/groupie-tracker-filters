let locations = document.getElementsByClassName("location")

for (let i = 0; i < locations.length; i++) {
    locations.item(i).innerHTML = locations.item(i).innerHTML.replaceAll("_", " ")
    locations.item(i).innerHTML = locations.item(i).innerHTML.replaceAll("-", ", ")


}
let searchInput = document.getElementsByName("q")[0]
searchButton = document.getElementById("search-button")

searchInput.addEventListener("input", function () {
    let q = searchInput.value.toLowerCase()
    if (q != "") {
        searchButton.href = `/search?q=${q}`
    } else{
        searchButton.removeAttribute("href")
    }
}, false)

artistsContainer = document.getElementById("container")
console.log(artistsContainer)
if (artistsContainer.childElementCount == 0) {
    p = document.createElement("p")
    p.style.color = "white"
    p.style.fontFamily = "Lato"
    p.style.fontSize = "x-large"
    text = document.createTextNode("No result found")
    p.appendChild(text)
    artistsContainer.appendChild(p)
    artistsContainer.style.justifyContent = "center"
    artistsContainer.style.alignItems = "center"

}


// =======
//     var yearRange = document.getElementById("year-range");
//     var selectedYearElement = document.getElementById("selected-year");

// function showDiv(divId) {
//     var myDiv = document.getElementById(divId);
//     myDiv.style.display = "block";
//   }

//   yearRange.addEventListener("input", function() {
//       var selectedYear = yearRange.value;
//       selectedYearElement.textContent = "Année sélectionnée : " + selectedYear;
      
//       // Utilisez la valeur sélectionnée pour filtrer les éléments affichés sur la page
//       // ou pour effectuer une autre action en fonction de l'année sélectionnée.
//   });

//   ============================================JS range==============================================================

const rangeInput1 = document.getElementById('myRange1');
const rangeInput2 = document.getElementById('myRange2');
const rangeInput3 = document.getElementById('myRange3');
const rangeInput4 = document.getElementById('myRange4');

const selectedValue1 = document.getElementById('selectedValue1');
const selectedValue2 = document.getElementById('selectedValue2');
const selectedValue3 = document.getElementById('selectedValue3');
const selectedValue4= document.getElementById('selectedValue4');

rangeInput1.addEventListener('input', function() {
  selectedValue1.textContent = rangeInput1.value;
});
rangeInput2.addEventListener('input', function() {
    selectedValue2.textContent = rangeInput2.value;
  });
  rangeInput3.addEventListener('input', function() {
    selectedValue3.textContent = rangeInput3.value;
  });
  rangeInput4.addEventListener('input', function() {
    selectedValue4.textContent = rangeInput4.value;
  });