let username = document.getElementById("Username");
let password = document.getElementById("Password");
let error=document.getElementById("Error");


async function login() {
	let res = await (await fetch("/api/user/login", {
		method: "POST",
		headers: {"Username": username.value, "Password": password.value}
	}));
	if (res["ok"]) {
		let JWT = await res.text();
		const date = new Date();
		date.setTime(Date.now() + (24 * 60 * 60 * 1000));
		document.cookie = "JWT=" + JWT + ";" + "Expires=" + date.toUTCString() + ";SameSite=strict;Path=/";
		window.location.href = "/items";
	} else {
		error.classList.remove("hidden");
		error.innerText = await res.text()
	}
}


document.getElementById("Submit").addEventListener("click", async function () {
	await login();
})

password.addEventListener("keydown", async function (e) {
	if (e.key === "Enter") {
		await login();
	}
})

