document.addEventListener("DOMContentLoaded", function (event) {
  const menuRulesNav = document.getElementById("rules-nav-title");
  const sectionMenu = document.getElementById("sections-menu");
  const sectionLinks = document.querySelectorAll(".section-link");
  const subsectionsExpand = document.querySelectorAll(".subsections-expand");
  const subsectionLinks = document.querySelectorAll(".subsection-link");
  const rulesNav = document.getElementById("rules-nav");

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
});
