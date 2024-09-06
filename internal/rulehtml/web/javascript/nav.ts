// closeNavRulesDropdown closes the menu opened by the "Rules" button in the middle of the top nav bar.
export function closeNavRulesDropdown() {
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
