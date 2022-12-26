;(function () {
	
	'use strict';

	// iPad and iPod detection	
	var isiPad = function(){
		return (navigator.platform.indexOf("iPad") != -1);
	};

	var isiPhone = function(){
	    return (
			(navigator.platform.indexOf("iPhone") != -1) || 
			(navigator.platform.indexOf("iPod") != -1)
	    );
	};

	// OffCanvass
	var offCanvass = function() {
		$('body').on('click', '.js-fh5co-menu-btn, .js-fh5co-offcanvass-close', function(){
			$('#fh5co-offcanvass').toggleClass('fh5co-awake');
		});
	};

	// Click outside of offcanvass
	var mobileMenuOutsideClick = function() {
		$(document).click(function (e) {
	    var container = $("#fh5co-offcanvass, .js-fh5co-menu-btn");
	    if (!container.is(e.target) && container.has(e.target).length === 0) {
	    	if ( $('#fh5co-offcanvass').hasClass('fh5co-awake') ) {
	    		$('#fh5co-offcanvass').removeClass('fh5co-awake');
	    	}
	    }
		});

		$(window).scroll(function(){
			if ( $(window).scrollTop() > 500 ) {
				if ( $('#fh5co-offcanvass').hasClass('fh5co-awake') ) {
		    		$('#fh5co-offcanvass').removeClass('fh5co-awake');
		    	}
	    	}
		});
	};

	// Magnific Popup
	
	var magnifPopup = function() {
		$('.image-popup').magnificPopup({
			type: 'image',
			removalDelay: 300,
			mainClass: 'mfp-with-zoom',
			titleSrc: 'title',
			gallery:{
				enabled:true
			},
			zoom: {
				enabled: true, // By default it's false, so don't forget to enable it

				duration: 300, // duration of the effect, in milliseconds
				easing: 'ease-in-out', // CSS transition easing function

				// The "opener" function should return the element from which popup will be zoomed in
				// and to which popup will be scaled down
				// By defailt it looks for an image tag:
				opener: function(openerElement) {
				// openerElement is the element on which popup was initialized, in this case its <a> tag
				// you don't need to add "opener" option if this code matches your needs, it's defailt one.
				return openerElement.is('img') ? openerElement : openerElement.find('img');
				}
			}
		});
	};



	var animateBoxWayPoint = function() {

		if ($('.animate-box').length > 0) {
			$('.animate-box').waypoint( function( direction ) {

				if( direction === 'down' && !$(this).hasClass('animated') ) {
					$(this.element).addClass('bounceIn animated');
				}

			} , { offset: '75%' } );
		}

	};


	//<div class="column size-1of2"> -- columnDiv
	// <div className="item"> -- photoDiv
	// 	<div className="animate-box"> -- photoInfoDiv
	// 		<a href="images/img_1.jpg" className="image-popup fh5co-board-img" title="描述xxxx"><img
	// 			src="images/img_1.jpg" alt=""></a> -- infoA infoImg
	// 	</div>
	// 	<div className="fh5co-item-title">标题</div> -- titleDiv
	// 	<div className="fh5co-desc">描述内容</div> -- descDiv
	// </div>
	//</div>

	var calPage = 0;
	var checkPage = 0;
	var columnValue;
	var C = 10; //滚动条距离底部的距离
	function initFirstPhotoList() {
		columnValue = document.getElementById("fh5co-board").getAttribute("data-columns");
		// 找到最外层的div
		var parentDiv = document.getElementById("fh5co-board");
		for(var idx=1; idx<=columnValue; idx++) {
			var columnDiv = document.createElement("div");
			columnDiv.setAttribute("id", "column_"+idx);
			columnDiv.setAttribute("class", "column size-1of"+columnValue);
			parentDiv.appendChild(columnDiv);
		}
		getPhotoList(calPage, 10, columnValue)
	}

	function initFirstPhotoList() {
		columnValue = document.getElementById("fh5co-board").getAttribute("data-columns");
		// 找到最外层的div
		var parentDiv = document.getElementById("fh5co-board");
		for(var idx=1; idx<=columnValue; idx++) {
			var columnDiv = document.createElement("div");
			columnDiv.setAttribute("id", "column_"+idx);
			columnDiv.setAttribute("class", "column size-1of"+columnValue);
			parentDiv.appendChild(columnDiv);
		}
		getPhotoList(calPage, 10, columnValue)
	}

	function getPhotoList(){
		calPage++
		if (calPage === checkPage+1 && calPage !== 1) {
			return
		}
		var protocol = window.location.protocol;
		var host = window.location.host;
		var url = protocol+"//"+host+"/photo-wall/photos?page="+calPage+"&page_size=5";
		$.getJSON(url, function(result){
			if (result.data.length === 0) {
				checkPage = calPage;
				return
			}
			
			for(var idx=0;idx<result.data.length;idx++) {
				var columnIdx = idx+1
				var pColumnDiv
				if (columnIdx%columnValue !== 0) {
					pColumnDiv = document.getElementById("column_"+columnIdx%columnValue)
				}else {
					pColumnDiv = document.getElementById("column_"+columnValue)
				}
				var photoUrl = result.data[idx].photo_url;
				var title = result.data[idx].photo_title;
				var describe = result.data[idx].describe;
				// 照片层div创建
				var photoDiv = document.createElement("div");
				// 设置属性
				photoDiv.setAttribute("class", "item");

				var photoInfoDiv = document.createElement("div");
				photoInfoDiv.setAttribute("class", "animate-box bounceIn animated");

				var infoA = document.createElement("a");
				infoA.setAttribute("href", photoUrl);
				infoA.setAttribute("class", "image-popup fh5co-board-img");
				infoA.setAttribute("title", describe);

				var infoImg = document.createElement("img");
				infoImg.setAttribute("src", photoUrl);
				infoImg.setAttribute("alt", "");

				var titleDiv = document.createElement("div")
				titleDiv.setAttribute("class", "fh5co-item-title");
				titleDiv.innerHTML = title;
				var descDiv = document.createElement("div")
				descDiv.setAttribute("class", "fh5co-desc");
				descDiv.innerHTML = describe;
				// 子div依赖
				pColumnDiv.appendChild(photoDiv);
				photoDiv.appendChild(photoInfoDiv);
				photoInfoDiv.appendChild(infoA);
				infoA.appendChild(infoImg);
				photoDiv.appendChild(titleDiv);
				photoDiv.appendChild(descDiv);
			}
		})
	}

	window.onscroll= function(){
		//文档内容实际高度（包括超出视窗的溢出部分）
		var scrollHeight = Math.max(document.documentElement.scrollHeight, document.body.scrollHeight);
		//滚动条滚动距离
		var scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop;
		//窗口可视范围高度
		var clientHeight = window.innerHeight || Math.min(document.documentElement.clientHeight,document.body.clientHeight);
		if(clientHeight + scrollTop + 0.5 >= scrollHeight){
			initFirstPhotoList();
		}
	}

	$(function(){
		initFirstPhotoList();
		magnifPopup();
		offCanvass();
		mobileMenuOutsideClick();
		animateBoxWayPoint();
	});


}());