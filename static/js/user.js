fetch('/users').then(function (response) {
    return response.json()
}).then(function (response) {
    console.log(window.location.href)
    const url = window.location.href.split('/')
    return response.filter(user => user.ID == url[url.length - 1])[0]
}).then(function (response) {
    console.log("response: " + response.Pseudo)
    document.getElementById("profilname").innerText = response.Pseudo
    document.getElementById("profilrank").innerText = (response.Level == "1") ? "Utilisateur" : (response.Level == "2") ? "Modérateur" : "Administrateur"
})