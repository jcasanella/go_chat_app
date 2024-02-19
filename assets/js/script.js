import { HttpError } from "./errors.js";

const URL = 'http://localhost:8080';

const isRequired = value => value === '' ? false : true;

const isBetween = (length, min, max) => length < min || length > max ? false : true;

const isPasswordSecure = (password) => {
    const re = new RegExp("^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})");
    return re.test(password);
};

const showError = (input, message) => {
    // get the form-field element
    const formField = input.parentElement;
    // add the error class
    formField.classList.remove('success');
    formField.classList.add('error');

    // show the error message
    const error = formField.querySelector('small');
    error.textContent = message;
};

const showSuccess = (input) => {
    // get the form-field element
    const formField = input.parentElement;

    // remove the error class
    formField.classList.remove('error');
    formField.classList.add('success');

    // hide the error message
    const error = formField.querySelector('small');
    error.textContent = '';
}

const usernameEl = document.querySelector('#username');
const checkUsername = () => {
    let valid = false;
    const min = 3,
        max = 25;
    const username = usernameEl.value.trim();

    if (!isRequired(username)) {
        showError(usernameEl, 'Username cannot be blank.');
    } else if (!isBetween(username.length, min, max)) {
        showError(usernameEl, `Username must be between ${min} and ${max} characters.`)
    } else {
        showSuccess(usernameEl);
        valid = true;
    }
    return valid;
}

const passwordEl = document.querySelector('#password');
const checkPassword = () => {
    let valid = false;

    const password = passwordEl.value.trim();

    if (!isRequired(password)) {
        showError(passwordEl, 'Password cannot be blank.');
    } else if (!isPasswordSecure(password)) {
        showError(passwordEl, 'Password must has at least 8 characters that include at least 1 lowercase character, 1 uppercase characters, 1 number, and 1 special character in (!@#$%^&*)');
    } else {
        showSuccess(passwordEl);
        valid = true;
    }

    return valid;
};

// Login call and validation
const signInFnc = async () => {
    const isUsernameValid = checkUsername(),
        isPasswordValid = checkPassword();

    const isFormValid = isUsernameValid && isPasswordValid;
    if (isFormValid) {

        const userName = usernameEl.value;
        const password = passwordEl.value;

        const user = {
            username: `${userName}`,
            password: `${password}`
        };


        try {
            const response = await fetch(`${URL}/api/login`, {
                method: 'POST',
                headers: {
                'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
            });

            if (!response.ok) {
                throw new HttpError(response.status, 'Error fetching data');
            }
            const resp = await response.json();
            console.log(`Resp: ${resp.token}`);
            sessionStorage.setItem("token", `${resp.token}`);

            // hide login
            let containerLogin = document.getElementById("container-login");
            containerLogin.style.display = "none";

            // show main app
            let containerMain = document.getElementById("main-container");
            containerMain.style.display = "block";

        } catch(err) {
            let modal = document.getElementById("modalDialog");
            let span = document.getElementsByClassName("close")[0];
            let errorBtn = document.getElementById("errorBtn");


            // When the user clicks on <span> (x), close the modal
            span.onclick = function() {
                modal.style.display = "none";
            }

            // When the user clicks anywhere outside of the modal, close it
            window.onclick = function(event) {
                if (event.target == modal) {
                    modal.style.display = "none";
                }
            }

            // When the user clicks on errorBtn, close the modal
            errorBtn.onclick = function() {
                modal.style.display = "none";
            }

            if (err instanceof HttpError) {
                modal.style.display = "block";
            }
        }
    }

};

document.getElementById("signInBtn").onclick = signInFnc;

// Starting signUp process
const signUpFnc = () => {
    console.log('Not implemented yet!!!')
}

document.getElementById("signUpBtn").onclick = signUpFnc;

window.onload  = function() {
    let token = sessionStorage.getItem("token");
    if (!token || token === '') { // Init
        // show login
        let containerLogin = document.getElementById("container-login");
        containerLogin.style.display = "block";

        // hide main app
        let containerMain = document.getElementById("main-container");
        containerMain.style.display = "none";
    } else {    // We have token - dont ask login again
        // hide login
        let containerLogin = document.getElementById("container-login");
        containerLogin.style.display = "none";

        // show main app
        let containerMain = document.getElementById("main-container");
        containerMain.style.display = "block";
    }
}