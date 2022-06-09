const menucontent = document.querySelector('#connectedmenucontent')

function hideMenu() {
  menucontent.classList.remove('open')
}

function showContent() {
  menucontent.classList.add('open')
}