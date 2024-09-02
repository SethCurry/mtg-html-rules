import uFuzzy from "./vendor/uFuzzy.esm";

// getContentsToSearch returns a tuple, where the first element is an array of HTML anchors
// and the second element is an array of strings to search through.
//
// The arrays are paired, so that the anchor at index i in the first array points to the
// contents of index i in the second array.  It's basically a dictionary decomposed into
// two arrays.
export function getContentsToSearch() {
  const searchHaystack: string[] = [];
  const searchIDs: string[] = [];

  document.querySelectorAll(".content").forEach((content) => {
    const parent = content.parentElement;
    if (!parent) {
      console.log("parent is null");
      return;
    }

    const ruleChild = parent.children[0] as HTMLParagraphElement;

    const ruleNumber = ruleChild.innerText;

    const ruleID = content.parentElement.id;

    const contentP = content as HTMLParagraphElement;
    searchHaystack.push(contentP.innerText);
    searchIDs.push(ruleID);
  });

  return [searchIDs, searchHaystack];
}

export function handleSearchInput(
  uf,
  searchHaystack: string[],
  searchIDs: string[]
) {
  const searchInput = document.getElementById("search-input");
  if (!searchInput) {
    console.log("search input is null");
    return;
  }

  const searchResultsDiv = document.getElementById("search-results");
  if (!searchResultsDiv) {
    console.log("search results div is null");
    return;
  }

  const searchInputInput = searchInput as HTMLInputElement;

  const searchTerm = searchInputInput.value.toLowerCase();
  let [idxs, info, order] = uf.search(searchHaystack, searchTerm);
  if (!order) {
    console.log("order is null");
    return;
  }

  searchResultsDiv.innerHTML = "";

  let domElems: HTMLElement[] = [];

  const mark = (part, matched): HTMLElement => {
    var el;

    if (matched) {
      el = document.createElement("mark");
      el.style.color = "black";
    } else {
      el = document.createElement("span");
      el.style.color = "white";
    }
    el.textContent = part;
    el.style.textDecoration = "none";
    return el;
  };

  // ordinarily we would just append to the accumulator, but because I can't
  // get uFuzzy to allow HTMLElement[] as the accumulator type, we have to use
  // an external accumulator.
  const append = (accum: HTMLElement[], part: HTMLElement): void => {
    accum.push(part);
  };

  for (let i = 0; i < order.length; i++) {
    let infoIdx = order[i];

    let matchEl = document.createElement("div");

    const parts = uFuzzy.highlight(
      searchHaystack[info.idx[infoIdx]],
      info.ranges[infoIdx],
      mark,
      [],
      append
    );

    parts.forEach((part) => {
      part.style.textDecoration = "none";
    });

    matchEl.append(...parts);
    matchEl.style.textDecoration = "none";

    let matchLink = document.createElement("a");
    matchLink.setAttribute("href", `#${searchIDs[info.idx[infoIdx]]}`);
    matchLink.style.textDecoration = "none";
    matchLink.append(matchEl);

    matchLink.addEventListener("click", () => {
      searchResultsDiv.style.display = "none";
    });

    matchLink.addEventListener("mouseover", () => {
      matchLink.style.backgroundColor = "black";
    });

    matchLink.addEventListener("mouseout", () => {
      matchLink.style.backgroundColor = "#46494C";
    });

    domElems.push(matchLink);
    const divider = document.createElement("hr");
    divider.style.color = "white";
    divider.style.width = "100%";
    domElems.push(divider);
  }

  searchResultsDiv.append(...domElems);

  searchResultsDiv.style.display = "flex";
}
