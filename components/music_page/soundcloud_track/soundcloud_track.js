document.addEventListener("DOMContentLoaded", () => {
  document.querySelectorAll(".sc-player-root").forEach((root) => {
    const iframe = root.querySelector(".sc-widget");
    const scUi = root.querySelector(".sc-ui");
    if (!iframe || !scUi) return;

    const widget = SC.Widget(iframe);
    const playBtn = scUi.querySelector(".sc-play");
    const pauseBtn = scUi.querySelector(".sc-pause");
    const seek = scUi.querySelector(".sc-seek");
    const elapsedEl = scUi.querySelector(".sc-elapsed");
    const durationEl = scUi.querySelector(".sc-duration");

    let duration = 0;

    const formatMs = (ms) => {
      const totalSeconds = Math.max(0, Math.floor(ms / 1000));
      const minutes = Math.floor(totalSeconds / 60);
      const seconds = totalSeconds % 60;
      return `${minutes}:${seconds.toString().padStart(2, "0")}`;
    };

    const updateTimes = (positionMs) => {
      if (!elapsedEl || !durationEl) return;
      const safeDuration = Math.max(duration, 0);
      const safePos = Math.min(Math.max(positionMs, 0), safeDuration || positionMs);
      elapsedEl.textContent = formatMs(safePos);
      durationEl.textContent = safeDuration ? formatMs(safeDuration) : "0:00";
    };

    function setSeekPercent(seekEl, percent) {
      if (!seekEl) return;
      const clamped = Math.max(0, Math.min(100, percent));
      seekEl.value = clamped;
      seekEl.style.setProperty("--sc-progress", `${clamped}%`);
    }

    widget.bind(SC.Widget.Events.READY, () => {
      widget.getDuration((d) => {
        duration = d;
        updateTimes(0);
      });
    });

    widget.bind(SC.Widget.Events.PLAY_PROGRESS, (e) => {
      if (duration > 0) {
        setSeekPercent(seek, (e.currentPosition / duration) * 100);
        updateTimes(e.currentPosition);
      }
    });

    if (playBtn) playBtn.onclick = () => widget.play();
    if (pauseBtn) pauseBtn.onclick = () => widget.pause();

    if (seek) {
      seek.oninput = (e) => {
        if (duration > 0) {
          const targetMs = (e.target.value / 100) * duration;
          widget.seekTo(targetMs);
          updateTimes(targetMs);
        }
      };
    }
  });
});
