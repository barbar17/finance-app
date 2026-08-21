tableSearchInit((query) => {
  state.search = query;

  loadUsers();
});

function renderUsers(users) {
  const tbody = document.querySelector("#table tbody");
  tbody.replaceChildren();

  if (users.length === 0) {
    handleEmptyTableData(tbody, 4, "User")

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

async function loadUsers() {
  const params = new URLSearchParams(state)

  const tbody = document.querySelector("#table tbody");
  tbody.replaceChildren();
  handleLoadingTableData(tbody, 4);

  const res = await fetch(`/api/users?${params}`)
  if (!res.ok) {
    throw new Error("Failed to load users")
  }

  const data = await res.json();

  renderUsers(data.data)
  renderPagination(data.pagination, loadUsers)
}

tableSortInit(loadUsers);
loadUsers();