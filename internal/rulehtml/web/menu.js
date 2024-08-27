document.addEventListener("DOMContentLoaded", function (event) {
  const menuRulesNav = document.getElementById("rules-nav-title");
  const sectionMenu = document.getElementById("sections-menu");
  const sectionLinks = document.querySelectorAll(".section-link");
  const subsectionsExpand = document.querySelectorAll(".subsections-expand");
  const subsectionLinks = document.querySelectorAll(".subsection-link");
  const rulesNav = document.getElementById("rules-nav");

  const searchHaystack = [];
  const searchIDs = [];

  document.querySelectorAll(".content").forEach((content) => {
    const parent = content.parentElement;
    const ruleNumber = parent.children[0].innerText;

    const ruleID = content.parentElement.id;
    searchHaystack.push(ruleNumber + " " + content.innerText);
    searchIDs.push(ruleID);
  });

  /*
  document.querySelectorAll(".rule").forEach((rule) => {
    const ruleID = rule.id;
    searchHaystack.push(rule.innerText);
    searchIDs.push(ruleID);
  });

  document.querySelectorAll(".subrule").forEach((subrule) => {
    const subruleID = subrule.id;
    searchHaystack.push(subrule.innerText);
    searchIDs.push(subruleID);
  });
  */

  function closeMenu() {
    sectionMenu.style.display = "none";
    sectionLinks.forEach((sectionLink) => {
      const sectionNum = sectionLink.getAttribute("sectionNum");
      const subsectionMenu = document.getElementById(
        `subsections-menu-${sectionNum}`
      );
      subsectionMenu.style.display = "none";
    });
  }

  sectionLinks.forEach((sectionLink) => {
    sectionLink.addEventListener("click", () => {
      closeMenu();
    });
  });

  subsectionLinks.forEach((subsectionLink) => {
    subsectionLink.addEventListener("click", () => {
      closeMenu();
    });
  });

  menuRulesNav.addEventListener("click", () => {
    sectionMenu.style.display = "block";
  });

  rulesNav.addEventListener("mouseleave", () => {
    closeMenu();
  });

  subsectionsExpand.forEach((subsectionExpander) => {
    const sectionNum = subsectionExpander.getAttribute("sectionNum");
    const subsectionMenu = document.getElementById(
      `subsections-menu-${sectionNum}`
    );

    subsectionExpander.addEventListener("click", () => {
      subsectionMenu.style.display = "block";
    });

    subsectionMenu.addEventListener("mouseleave", () => {
      subsectionMenu.style.display = "none";
    });
  });

  let opts = {};
  let uf = new uFuzzy(opts);

  const searchResultsDiv = document.getElementById("search-results");
  const searchInput = document.getElementById("search-input");

  searchInput.addEventListener("input", () => {
    const searchTerm = searchInput.value.toLowerCase();
    let [idxs, info, order] = uf.search(searchHaystack, searchTerm);

    searchResultsDiv.innerHTML = "";

    let domElems = [];

    const mark = (part, matched) => {
      let el = matched
        ? document.createElement("mark")
        : document.createElement("span");
      el.textContent = part;
      el.style.textDecoration = "none";
      el.style.color = "white";
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

      /*
      let resultLink = document.createElement("a");
      resultLink.setAttribute("href", `#${searchIDs[idxs[i][0]]}`);
      resultLink.textContent = searchHaystack[info.idx[order[i]]];

      searchResultsDiv.append(resultLink);
      */
    }

    searchResultsDiv.append(...domElems);

    searchResultsDiv.style.display = "flex";
  });
});
