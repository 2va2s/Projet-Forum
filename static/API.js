fetch("/users").then(function (response) {
    return response.json()
}).then(function (response) {
    console.log(response)
}).catch(function (err) {
    console.log(err)
})