function encodeImageFileAsURL(inputFileToLoad) {

    var filesSelected = document.getElementById("inputFileToLoad").files[0];
    if (filesSelected.length > 0) {
        var fileToLoad = filesSelected[0];

        var fileReader = new FileReader();

        fileReader.onload = function(fileLoadedEvent) {
            var srcData = fileLoadedEvent.target.result; // <--- data: base64
            console.log("srcData: " + srcData);
            var newImage = document.createElement('img');
            newImage.src = srcData;

            document.getElementById("imgTest").innerHTML = newImage.outerHTML;
            console.log("Converted Base64 version is " + document.getElementById("imgTest").innerHTML);
        }
        fileReader.readAsDataURL(fileToLoad);
    }
}