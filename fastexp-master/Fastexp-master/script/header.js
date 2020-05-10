let img = document.querySelector(".second");
let block = document.querySelector(".retractable");
let def = document.querySelector(".default");
let profile = document.querySelector(".name");
let logo = document.querySelector(".logo");

logo.onclick = () => {
  window.location = "themes.html";
};

def.children[0].onclick = () => {
  window.location = "profile.html";
};

block.children[0].onclick = () => {
  window.location = "profile.html";
};

img.onclick = () => {
  block.classList.toggle("hidden");
};

def.children[1].onclick = () => {
  window.location = "auth.html";
};

block.children[1].onclick = () => {
  window.location = "auth.html";
};
