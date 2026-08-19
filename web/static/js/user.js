function searchUser(query) {
  state.search = query;

  loadUsers();
}

tableSearchInit(searchUser);

function renderUsers(users) {
  const tbody = document.querySelector("#table tbody");
  tbody.replaceChildren();

  if (users.length === 0) {
    const row = document.createElement("tr");
    const cell = document.createElement("td");

    cell.colSpan = 4;
    cell.textContent = "No users found.";
    cell.style.textAlign = "center";

    row.appendChild(cell);
    tbody.appendChild(row);

    return;
  }

  users.forEach((user, index) => {
    const row = document.createElement("tr");

    const no = document.createElement("td");
    const username = document.createElement("td");
    const name = document.createElement("td");
    const createdAt = document.createElement("td");

    no.textContent = String(index + 1);
    username.textContent = user.Username;
    name.textContent = user.Name;
    createdAt.textContent = formatDate(user.CreatedAt);

    row.appendChild(no);
    row.appendChild(username);
    row.appendChild(name);
    row.appendChild(createdAt);

    tbody.appendChild(row);
  })
}

function renderPagination(pagination) {
  const container = document.querySelector("#table-pagination");
  container.replaceChildren();

  const { page, totalPages } = pagination;

  if (totalPages <= 1) {
    return;
  }

  const prev = document.createElement("button");
  prev.textContent = "Prev";
  prev.disabled = page === 1;
  prev.addEventListener("click", () => {
    state.page = page - 1;
    loadUsers();
  });

  container.appendChild(prev);

  for (let i = 1; i <= totalPages; i++) {
    const button = document.createElement("button");
    button.textContent = i;
    button.disabled = i === page;
    button.addEventListener("click", () => {
      state.page = i;
      loadUsers();
    });

    container.appendChild(button);
  }

  const next = document.createElement("button");
  next.textContent = "Next";
  next.disabled = page === totalPages;
  next.addEventListener("click", () => {
    state.page = page + 1;
    loadUsers();
  });

  container.appendChild(next);
}

async function loadUsers() {
  const params = new URLSearchParams(state)

  const res = await fetch(`/api/users?${params}`)
  if (!res.ok) {
    throw new Error("Failed to load users")
  }

  const data = await res.json();

  renderUsers(data.data)
  renderPagination(data.pagination)
}

loadUsers()