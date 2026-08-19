let searchTimeout;
const state = {
  page: 1,
  limit: 5,
  search: "",
  sort: "created_at",
  order: "desc",
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