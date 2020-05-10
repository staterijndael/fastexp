// let firstName = document.querySelector("#firstname");
// let lastName = document.querySelector("#lastname");
let email = document.querySelector("#email");
let pass = document.querySelector("#pass");
let passAgain = document.querySelector("#pass-again");
let btn = document.querySelector("#submit");
let auth = document.querySelector("#auth");

auth.onclick = function () {
  window.location = "auth.html";
};

btn.onclick = () => {
  if (pass.value === passAgain.value) {
    const regReq = async () => {
      let myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");

      let raw = JSON.stringify({
        email: email.value,
        password: pass.value,
      });

      let requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
      };

      try {
        let v = await fetch("http://localhost:8080/users", requestOptions);
        if (v.ok) {
          window.location = "themes.html";
        }
      } catch (err) {
        console.log(err);
      }
    };

    regReq();
  }
};
