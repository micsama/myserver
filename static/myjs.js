// 1rem -- 16px

// 监听浏览器,针对不同分辨率计算font-size
// 然后根据设计稿比如尺寸是640尺寸 rem = 设计稿的字体大小 / 100 ; 16px —> 0.16rem
// dpr 设备像素比

(function (doc, win) {
	var docEl = doc.documentElement,
		resizeEvt = "orientationchange" in window ? "orientationchange" : "resize",
		recalc = function () {
			var clientWidth = docEl.clientWidth;
			if (!clientWidth) return;
			if (clientWidth <= 320) {
				docEl.style.fontSize = "350px";
			} else if (clientWidth >= 640) {
				docEl.style.fontSize = "200px";
			} else {
				docEl.style.fontSize = 300 * (clientWidth / 640) + "px";
			}
		};
	if (!doc.addEventListener) return;
	win.addEventListener(resizeEvt, recalc, false);
	doc.addEventListener("DOMContentLoaded", recalc, false);
})(document, window);

// 按照16 px=0.16rem 来设置字体和宽高，问题基本就能解决来！！！
