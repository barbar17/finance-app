function searchTransaction(query) {
  state.search = query;

  loadTransactions();
}

tableSearchInit(searchTransaction);

function renderTransactions(transactions) {
  const tbody = document.querySelector("#table tbody");
  tbody.replaceChildren();

  if (transactions.length === 0) {
    const row = document.createElement("tr");
    const cell = document.createElement("td");

    cell.colSpan = 6;
    cell.textContent = "No transaction found.";
    cell.style.textAlign = "center";

    row.appendChild(cell);
    tbody.appendChild(row);

    return;
  }

  transactions.forEach((transaction, index) => {
    const row = document.createElement("tr");

    const no = document.createElement("td");
    const name = document.createElement("td");
    const amount = document.createElement("td");
    const desc = document.createElement("td");
    const category = document.createElement("td");
    const createdAt = document.createElement("td");

    no.textContent = String(index + 1);
    name.textContent = transaction.Name;
    amount.textContent = transaction.Amount;
    desc.textContent = transaction.Desc;
    category.textContent = transaction.Category;
    createdAt.textContent = formatDate(transaction.CreatedAt);

    row.appendChild(no);
    row.appendChild(name);
    row.appendChild(amount);
    row.appendChild(desc);
    row.appendChild(category);
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
    state.page = Number(page) - 1;
    loadTransactions();
  });

  container.appendChild(prev);

  for (let i = 1; i <= totalPages; i++) {
    const button = document.createElement("button");
    button.textContent = i;
    button.disabled = i === page;
    button.addEventListener("click", () => {
      state.page = i;
      loadTransactions();
    });

    container.appendChild(button);
  }

  const next = document.createElement("button");
  next.textContent = "Next";
  next.disabled = page === totalPages;
  next.addEventListener("click", () => {
    state.page = Number(page) + 1;
    loadTransactions();
  });

  container.appendChild(next);
}

async function loadTransactions() {
  const params = new URLSearchParams(state)

  const res = await fetch(`/api/transactions?${params}`)
  if (!res.ok) {
    throw new Error("Failed to load transactions")
  }

  const data = await res.json();

  renderTransactions(data.data)
  renderPagination(data.pagination)
}

loadTransactions()