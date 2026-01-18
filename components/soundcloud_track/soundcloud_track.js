document.addEventListener("DOMContentLoaded", () => {
  const iframe = document.querySelector(".sc-widget");
  const widget = SC.Widget(iframe);

  const scUi = document.querySelector(".sc-ui");
  const playBtn = scUi.querySelector(".sc-play");
  const pauseBtn = scUi.querySelector(".sc-pause");
  const seek = scUi.querySelector(".sc-seek");

  let duration = 0;


  function setSeekPercent(seekEl, percent) {
    const clamped = Math.max(0, Math.min(100, percent));
    seekEl.value = clamped;
    seekEl.style.setProperty("--sc-progress", `${clamped}%`);
  }

  widget.bind(SC.Widget.Events.READY, () => {
    widget.getDuration(d => duration = d);
  });

  widget.bind(SC.Widget.Events.PLAY_PROGRESS, (e) => {
    if (duration > 0) {
      setSeekPercent(seek, (e.currentPosition / duration) * 100);
    }
  });

  playBtn.onclick = () => widget.play();
  pauseBtn.onclick = () => widget.pause();

  seek.oninput = (e) => {
    if (duration > 0) {
      widget.seekTo((e.target.value / 100) * duration);
    }
  };
});
