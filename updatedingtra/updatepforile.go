package updatedingtra


import (
	"html/template"
	// "rundb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (let *UpdateStruct ) ProfileForm (item primitive.M) {
	let.ParseFields(item)
	let.Details = template.HTML(
	// <html>
	// 	<head>
	// 	    <title>Update profile</title>
	// 	    <link rel="icon" href="/jskl899nxhsjas/" type="image/x-icon">
	// 	    <meta charset="UTF-8">
	// 	    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	// 	    <link rel="stylesheet" href="/dingtra-web-assets/css/cssupdte001/updt090.css" />
	// 	    <script src="/dingtra-web-assets/js/feather/feather.js"></script>
	// 		<script src="/dingtra-web-assets/js/jsquery/jquery.js"></script>
	// 		<script src="/dingtra-web-assets/js/jsupres/jsupres001.js"></script>
	// 	</head>
	// 	<body>
		`
		
	    <div class="updatexr0980"><div class="updatexr0980-chld">
		   <div id="ntfyxr0k87cncl"><div style="pointer-events:none;display:flex;justify-content:end;padding:10px 10px 10px 0px;">
		   <button><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="black" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="feather feather-x" color="black"><line x1="18" y1="6" x2="6" y2="18"></line><line x1="6" y1="6" x2="18" y2="18"></line></svg></button> 
		   </div></div>
		   <div id="errsklop90"></div>
		    <form enctype='multipart/form-data' name="updtfilexr09">
				<div class="updatexr0980-img">
				   <input id="imagefiler" style="display:none;" type="file" name="file" onchange="showPreview(event);">
				    <div id="img-trash09x" class="imgtrash090">
					    <div style="pointer-events:none;">
				            <span> <i data-feather="trash" color="black"></i> </span>
						</div>
					</div>
				   <div id="filexr09780"><label class="xr09bklabel" for="imagefiler">`+let.Pic(item)+`</label></div>
				</div>

				<div id="imagetrsh09xr" style="display:none;padding:10px 0px 10px 0px;width:fit-content;margin:auto;">
					<div class="imgtrashmore">
						<div id="imgtrsh090" class="imgtrashmore090"> 
							<div style="pointer-events:none;"><span>Remove image</span></div>
						</div>
						<div id="imgtrsh091" style="border:unset;" class="imgtrashmore090"> 
							<span style="pointer-events:none;">Cancel</span>
						</div>
					</div>
				</div>

					
	
				<div class="updatexr0980-box">
					<div class="jsr890">
					   <div class="jsr891"><span>Fullname</span></div>
					   <div class="jsr892"><span> <input name="fullname" type="text" value="`+let.Fullname+`"/> </span></div>
					</div>
	
					<div class="jsr890">
					   <div class="jsr891"><span>Username</span></div>
					   <div class="jsr892"><span> <input name="username" type="text" value="`+let.Username+`"/> </span></div>
					</div>
	
					<div class="jsr890">
					   <div class="jsr891"><span>Website</span></div>
					   <div class="jsr892"><span> <input name="website" type="text" value="`+let.Website+`"/> </span></div>
					</div>

					<div class="jsr890">
					   <div class="jsr891"><span>Work</span></div>
					   <div class="jsr892"><span> <input name="work" type="text" value="`+let.Work+`"/> </span></div>
					</div>
	
					<div class="jsr890">
					   <div class="jsr891"><span>Location</span></div>
					   <div class="jsr892"><span> <input name="location" type="text" value="`+let.Location+`"/> </span></div>
					</div>
				</div>
			</form>

			<div class="jsr890">
				<div id="updatexr090"><div style="pointer-events:none;" class="jsr891 btnxr0p90"><span> <button>Update profile</button> </span></div></div>
			</div>
	
		</div></div>

		<script>
        feather.replace()
        </script>
	`)
}


func (let *UpdateStruct ) Pic(item primitive.M) string {
	var res string
	if item["pic"] != nil && len(item["pic"].(string)) > 0 {
		if item["pic"].(string)[0:1] == "#" {
			res = "<img src='"+item["pic"].(string)[1:]+"' />"
		}else{
			res  = "<img src='/dingtra-web-assets/images/profile/"+item["pic"].(string)+"' />"
		}
	}else{
		res  = "<img src='/dingtra-web-assets/svg/pr.png' />"
	}

	return res
	
}

func (let *UpdateStruct ) ParseFields (item primitive.M) {

	let.Username = item["username"].(string)

	if item["fullname"] != nil && len(item["fullname"].(string)) > 0 {

		let.Fullname = item["fullname"].(string)
	}

	if item["location"] != nil && len(item["location"].(string)) > 0 {

		let.Location = item["location"].(string)
	}

	if item["about"] != nil && len(item["about"].(string)) > 0 {

		let.Bio = item["about"].(string)
	}

	if item["website"] != nil && len(item["website"].(string)) > 0 {

		let.Website = item["website"].(string)
	}

	if item["work"] != nil && len(item["work"].(string)) > 0 {

		let.Work = item["work"].(string)
	}



}
