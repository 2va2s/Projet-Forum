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
                return
            }
            errorP = document.getElementById("loginError")
            errorP.innerText = "Identifiants incorrects"
        }
    }
    ).catch(function (err) {
        console.log(err)
    })
}

// const checkRegisterForm = () => {
//     fetch("/users").then(function (response) {
//         return response.json()
//     }).then(function (response) {
//         // let pseudo = document.getElementById('pseudo2')
//         // let mdp1 = document.getElementById('password2')
//         // let mdp2 = document.getElementById('password22')
//         // let email = document.getElementById('email2')
//         // let telephone = document.getElementById('telephone')
//         // if (pseudo.value.lenght < 4) raiseRegisterError("Le mot de passe doit contenir au mois 4 caractères")
//         // if (pseudo.value.includes(" ") < 4) raiseRegisterError("Le pseudo ne doit pas contenir d'espaces (utiliser un '_' ou un '-'")
//         // if (mdp1.value != mdp2.value) raiseRegisterError()




//         for (let i in response) {
//             if (response[i].Pseudo == email.value && response[i].Password == mdp.value) {
//                 document.getElementById('subscribe').submit()
//             } else {
//                 console.log("true")
//             }
//         }
//     })
// }