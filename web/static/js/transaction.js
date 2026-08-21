tableSearchInit((query) => {
  state.search = query;

  loadTransactions();
});

function renderTransactions(transactions) {
  const tbody = document.querySelector("#table tbody");
  tbody.replaceChildren();

  if (transactions.length === 0) {
    handleEmptyTableData(tbody, 6, "Transaction")

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

    const rootStyles = getComputedStyle(document.documentElement);

    no.textContent = "TX-" + transaction.ID;
    name.textContent = transaction.Name;
    amount.textContent = transaction.Amount;
    amount.style.backgroundColor = Number(transaction.Amount) >= 0
      ? rootStyles.getPropertyValue("--primary")
      : rootStyles.getPropertyValue("--danger");
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

async function loadTransactions() {
  const params = new URLSearchParams(state)

  const tbody = document.querySelector("#table tbody");
  tbody.replaceChildren();
  handleLoadingTableData(tbody, 6);

  const res = await fetch(`/api/transactions?${params}`)
  if (!res.ok) {
    throw new Error("Failed to load transactions")
  }

  const data = await res.json();

  renderTransactions(data.data)
  renderPagination(data.pagination, loadTransactions)
}

tableSortInit(loadTransactions);
loadTransactions()