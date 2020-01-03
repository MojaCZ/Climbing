window.onload = function(evt){
  setPageDim()
  window.addEventListener("resize",function(event){setPageDim()})
  console.log("HELLO");

};


function setPageDim(){
  var menuDiv = document.getElementById("menu")
  var story = document.getElementById("themeImageDiv")
  var storyInfo = document.getElementById("mainStoryInfo")
  var aboutWeb = document.getElementById("mainPageContent")
  // menu offset
  var menuOffset = menuDiv.offsetHeight

  if (storyInfo == null) {
    storyInfoHeight = 0
    menuOffset = menuOffset -100
  } else {
    storyInfoHeight = storyInfo.offsetHeight
    storyInfo.style.marginBottom = -storyInfoHeight;
  }
  var docWidth = window.innerWidth
  var docHeight = window.innerHeight


  var w = docWidth - 1130 // breaking point, where triangles fits (ratio is 1) is around 1125, I need to calculate triangle from this point
  var h = -docHeight
  menuOffset = 88/325*w+menuOffset // calculate ratio (numbers were tested and calculated as triangle ratios)..
  heightOffset = 0.5*h + 460
  menuOffset = menuOffset + heightOffset
  console.log(heightOffset, menuOffset);
  if (docWidth > 1600) {menuOffset = menuOffset-15}  // for wide monitors
  if (docWidth < 770) {menuOffset = menuOffset-10}  // for tablets
  var storyOffset = story.offsetHeight

  story.style.top = menuOffset.toString() + "px"
  aboutWeb.style.top = (storyOffset-storyInfoHeight+40).toString() + "px"
}
