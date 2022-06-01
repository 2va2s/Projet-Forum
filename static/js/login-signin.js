const toLogin = () => {
    // document.getElementById("login").style.transition = "margin 1000ms ease-in-out"
    // document.getElementById("register").style.transition = "margin 1000ms ease-in-out"
    document.getElementById("register").style.marginLeft = "-680%"
    // document.getElementById("register").style.
    document.getElementById("login").style.marginLeft = "0"
}

const toRegister = () => {
    // document.getElementById("register").style.transition = "margin 1000ms ease-in-out"
    // document.getElementById("login").style.transition = "margin 1000ms ease-in-out"
    document.getElementById("login").style.marginLeft = "-680%"
    document.getElementById("register").style.marginLeft = "0"
}

const checkLoginForm = () => {
    fetch("/users").then(function (response) {
        return response.json()
    }).then(function (response) {
        for (let i in response) {
            let email = document.getElementById('email1')
            let mdp = document.getElementById('mdp1')
            if (response[i].Pseudo == email.value && response[i].Password == mdp.value) {
                document.getElementById('subscribe').submit()
            } else {
                // %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
                // §§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§
                // CREER MESSAGE QUI DIT IDENTIFIANTS INCORRECT
                // §§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§§
                // %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
            }
        }
    }
    ).catch(function (err) {
        console.log(err)
    })
}

const checkRegisterForm = () => {
    fetch("/users").then(function (response) {
        return response.json()
    }).then(function (response) {
        
    }
}