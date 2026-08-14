document
  .getElementById("login-form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();

    const msg = document.getElementById("msg");
    try {
      const res = await fetch("/login", {
        method: "POST",
        body: new FormData(e.target),
      })

      const data = await res.json();
      if (res.ok) {
        window.location.href = data.redirect
      } else {
        msg.textContent = data.msg
      }

    } catch (err) {
      msg.textContent = "Server Unavailable";
    }
  })