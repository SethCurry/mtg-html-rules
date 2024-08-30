// getContentsToSearch returns a tuple, where the first element is an array of HTML anchors
// and the second element is an array of strings to search through.
//
// The arrays are paired, so that the anchor at index i in the first array points to the
// contents of index i in the second array.  It's basically a dictionary decomposed into
// two arrays.
function getContentsToSearch() {
  const searchHaystack = [];
  const searchIDs = [];

  document.querySelectorAll(".content").forEach((content) => {
    const parent = content.parentElement;
    const ruleNumber = parent.children[0].innerText;

    const ruleID = content.parentElement.id;
    searchHaystack.push(ruleNumber + " " + content.innerText);
    searchIDs.push(ruleID);
  });

  return [searchIDs, searchHaystack];
}

// closeRulesMenu closes the menu opened by the "Rules" button in the middle of the top nav bar.
function closeRulesMenu() {
  const sectionMenu = document.getElementById("sections-menu");
  const sectionLinks = document.querySelectorAll(".section-link");

  sectionMenu.style.display = "none";
  sectionLinks.forEach((sectionLink) => {
    const sectionNum = sectionLink.getAttribute("sectionNum");
    const subsectionMenu = document.getElementById(
      `subsections-menu-${sectionNum}`
    );
    subsectionMenu.style.display = "none";
  });
}

function handleSearchInput(uf, searchHaystack, searchIDs) {
  const searchInput = document.getElementById("search-input");
  const searchResultsDiv = document.getElementById("search-results");

  const searchTerm = searchInput.value.toLowerCase();
  let [idxs, info, order] = uf.search(searchHaystack, searchTerm);

  searchResultsDiv.innerHTML = "";

  let domElems = [];

  const mark = (part, matched) => {
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

  const append = (accum, part) => {
    accum.push(part);
  };

  for (let i = 0; i < order.length; i++) {
    let infoIdx = order[i];

    let matchEl = document.createElement("div");

    let parts = uFuzzy.highlight(
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

document.addEventListener("DOMContentLoaded", function (event) {
  const menuRulesDropdownButton = document.getElementById(
    "rules-dropdown-button"
  );
  const sectionMenu = document.getElementById("sections-menu");
  const sectionLinks = document.querySelectorAll(".section-link-name");
  const subsectionsExpand = document.querySelectorAll(".subsections-expand");
  const subsectionLinks = document.querySelectorAll(".subsection-link");
  const rulesNav = document.getElementById("rules-nav");

  const [searchIDs, searchHaystack] = getContentsToSearch();

  sectionLinks.forEach((sectionLink) => {
    sectionLink.addEventListener("click", () => {
      const ruleNum = sectionLink.getAttribute("sectionNum");
      closeRulesMenu();
      window.location.hash = `#section-${ruleNum}`;
    });
  });

  subsectionLinks.forEach((subsectionLink) => {
    subsectionLink.addEventListener("click", () => {
      const subsectionNum = subsectionLink.getAttribute("subsection");
      closeRulesMenu();
      window.location.hash = `#rule-${subsectionNum}`;
    });
  });

  menuRulesDropdownButton.addEventListener("click", () => {
    sectionMenu.style.display = "flex";
  });

  rulesNav.addEventListener("mouseleave", () => {
    closeRulesMenu();
  });

  subsectionsExpand.forEach((subsectionExpander) => {
    const sectionNum = subsectionExpander.getAttribute("sectionNum");
    const subsectionMenu = document.getElementById(
      `subsections-menu-${sectionNum}`
    );

    subsectionExpander.addEventListener("click", () => {
      if (subsectionMenu.style.display === "block") {
        subsectionMenu.style.display = "none";
      } else {
        subsectionMenu.style.display = "block";
      }
    });
  });

  let opts = {};
  let uf = new uFuzzy(opts);

  const searchInput = document.getElementById("search-input");

  searchInput.addEventListener("input", () => {
    handleSearchInput(uf, searchHaystack, searchIDs);
  });
});
