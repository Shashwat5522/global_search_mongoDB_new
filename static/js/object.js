$(document).ready(async function(){
    console.log("Script loaded")
    const urlParams = new URLSearchParams(window.location.search);
    const myParam = urlParams.get('objeId');
    let object=await fetch("http://localhost:8080/object",{headers:{"object-Id":myParam}})
    let jsonObject=await object.json()
    $(".titleClass").html(jsonObject.Type)
    $(".card-title").html(jsonObject.Title)
    $(".card-text").html(jsonObject.Description+"<br>"+jsonObject.Tags)
    $(".card-footer").html(jsonObject.Visible)

})