let searchTimeout;
const state = {
  page: 1,
  limit: 5,
  search: "",
  sort: "created_at",
  order: "desc",
}

function tableSortInit(loadDataFunction) {
  document.querySelectorAll("#table th[data-sort]").forEach((th) => {
    th.addEventListener("click", () => {
      const prevSortColumn = document.querySelector(`#table th[data-sort="${state.sort}"]`);
      prevSortColumn.querySelector("span").textContent = "";

      const sortIcon = th.querySelector("span");
      if (state.sort === th.dataset.sort) {
        state.order = state.order === "desc" ? "asc" : "desc"
        sortIcon.innerHTML = state.order === "desc"
          ? '<i class="fa-solid fa-arrow-down"></i>'
          : '<i class="fa-solid fa-arrow-up"></i>';
      } else {
        state.sort = th.dataset.sort
        state.order = "desc"
        sortIcon.innerHTML = '<i class="fa-solid fa-arrow-down"></i>';
      }

      loadDataFunction();
    })
  })
}

const tableSearch = document.getElementById("table-search");
function tableSearchInit(searchFunction) {
  tableSearch.addEventListener("input", (e) => {
    clearTimeout(searchTimeout);

    searchTimeout = setTimeout(() => {
      searchFunction(e.target.value);
    }, 300);
  })
}

function handleEmptyTableData(tBody, columnLength, dataName) {
  const row = document.createElement("tr");
  const cell = document.createElement("td");

  cell.colSpan = columnLength;
  cell.textContent = `No ${dataName} found.`;
  cell.style.textAlign = "center";

  row.appendChild(cell);
  tBody.appendChild(row);
}

function handleLoadingTableData(tBody, columnLength) {
  const row = document.createElement("tr");
  const cell = document.createElement("td");

  cell.colSpan = columnLength;
  cell.textContent = "Loading...";
  cell.style.textAlign = "center";

  row.appendChild(cell);
  tBody.appendChild(row);
}

function renderPagination(pagination, loadDataFunction) {
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
    loadDataFunction();
  });

  container.appendChild(prev);

  for (let i = 1; i <= totalPages; i++) {
    const button = document.createElement("button");
    button.textContent = i;
    button.disabled = i === page;
    button.addEventListener("click", () => {
      state.page = i;
      loadDataFunction();
    });

    container.appendChild(button);
  }

  const next = document.createElement("button");
  next.textContent = "Next";
  next.disabled = page === totalPages;
  next.addEventListener("click", () => {
    state.page = Number(page) + 1;
    loadDataFunction();
  });

  container.appendChild(next);
}