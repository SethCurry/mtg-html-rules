import {
  getContentsToSearch,
  handleSearchInput,
  hideSearchResults,
} from "./search";
import uFuzzy from "./vendor/uFuzzy.esm";

// closeRulesMenu closes the menu opened by the "Rules" button in the middle of the top nav bar.
function closeRulesMenu() {
  const sectionMenu = document.getElementById("sections-menu");
  if (!sectionMenu) {
    console.log("sections menu is null");
    return;
  }

  const sectionLinks = document.querySelectorAll(".section-link");

  sectionMenu.style.display = "none";
  sectionLinks.forEach((sectionLink) => {
    const sectionNum = sectionLink.getAttribute("sectionNum");
    const subsectionMenu = document.getElementById(
      `subsections-menu-${sectionNum}`
    );

    if (!subsectionMenu) {
      console.log("subsections menu is null");
      return;
    }
    subsectionMenu.style.display = "none";
  });
}

document.addEventListener("DOMContentLoaded", function (event) {
  const menuRulesDropdownButton = document.getElementById(
    "rules-dropdown-button"
  );
  if (!menuRulesDropdownButton) {
    console.log("menu rules dropdown button is null");
    return;
  }

  const sectionMenu = document.getElementById("sections-menu");
  if (!sectionMenu) {
    console.log("sections menu is null");
    return;
  }

  const sectionLinks = document.querySelectorAll(".section-link-name");
  const subsectionsExpand = document.querySelectorAll(".subsections-expand");
  const subsectionLinks = document.querySelectorAll(".subsection-link");

  const rulesNav = document.getElementById("rules-nav");
  if (!rulesNav) {
    console.log("rules nav is null");
    return;
  }

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
    if (!subsectionMenu) {
      console.log("subsection menu is null");
      return;
    }

    subsectionExpander.addEventListener("click", () => {
      if (subsectionMenu.style.display === "block") {
        subsectionMenu.style.display = "none";
      } else {
        subsectionMenu.style.display = "block";
      }
    });
  });

  let opts = {};
  let uf = uFuzzy(opts);

  const searchInput = document.getElementById("search-input");
  if (!searchInput) {
    console.log("search input is null");
    return;
  }

  searchInput.addEventListener("input", () => {
    handleSearchInput(uf, searchHaystack, searchIDs);
  });
  searchInput.addEventListener("focusout", () => {
    hideSearchResults();
  });
});
