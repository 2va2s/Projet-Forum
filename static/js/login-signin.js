const signUpButton = document.getElementById('signUp');
const signInButton = document.getElementById('signIn');
const container = document.getElementById('container');
const logBlocker = document.getElementById('logblocker')
const signIn = document.getElementById('menusignin')
const signUp = document.getElementById('menusignup')

signUpButton.addEventListener('click', () => {
	container.classList.add("right-panel-active");
});

signInButton.addEventListener('click', () => {
	container.classList.remove("right-panel-active");
});

logBlocker.addEventListener('click', () => {
	document.getElementById('container').style.display = 'none'
    logBlocker.style.display = 'none'
    document.body.style.overflow = 'auto'
});

signUp.addEventListener('click', () => {
	container.style.display = 'block'
    logBlocker.style.display = 'block'
    container.classList.add("right-panel-active");
    document.body.style.overflow = 'hidden'
});

signIn.addEventListener('click', () => {
	container.style.display = 'block'
    logBlocker.style.display = 'block'
    document.body.style.overflow = 'hidden'
});

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