fetch('/users').then(function (response) {
    return response.json()
}).then(function (response) {
    console.log(window.location.href)
    const url = window.location.href.split('/')
    return response.filter(user => user.ID == url[url.length - 1])
}).then(function (response) {
    
})
