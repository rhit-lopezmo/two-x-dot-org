(() => {
  if (window.__pageNavInit) return;
  window.__pageNavInit = true;

  /** @type {DOMParser} */
  const domParser = new DOMParser();

  /** @type {HTMLElement | null} */
  const content = document.querySelector("#content");

  /** @type {NodeListOf<HTMLAnchorElement>} */
  const links = document.querySelectorAll(".page-link");

  links.forEach((link) => {
    link.addEventListener("click", async (event) => {
      event.preventDefault();

      try {
        const res = await fetch(link.href);
        if (!res.ok) throw new Error(`HTTP ${res.status}`);

        const html = await res.text();
        const doc = domParser.parseFromString(html, "text/html");
        const newContent = doc.querySelector("#content");
        if (!newContent || !content) throw new Error("content missing");

        content.innerHTML = newContent.innerHTML;
        history.pushState({}, "", link.href);

        if (typeof window.initSoundcloudPlayers === "function") {
          window.initSoundcloudPlayers(content);
        }
      } catch (error) {
        console.error("Navigation failed, falling back to full load:", error);
        window.location.href = link.href;
      }
    });
  });

  window.addEventListener("popstate", () => {
    window.location.reload();
  });
})();
