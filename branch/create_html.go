package branch

import (
	"net/http"
	"html/template"
	"strings"
)

func (let *BranchStruct ) HTML (r *http.Request) {

	brncharr := []string{}


	// parent
	brncharr = append(brncharr, "<div class='new-branch-div-class'><div class='new-branch-div-class-ch'>")

	brncharr = append(brncharr, "<div style='display:none;'><div class='duplxk090'><textarea id='duplxjss98'></textarea></div></div>")

	// success
	brncharr = append(brncharr, "<div id='br909resxy'></div>")
	// success

	brncharr = append(brncharr, "<div class='new-branch-div-class-title'><span>Start a new branch</span></div>")
	brncharr = append(brncharr, "<div class='new-branch-div-class-details'><span>A branch is a workspace for your files/team . Record, manage and distribute</span></div>")

	// form
	brncharr = append(brncharr, "<form name='frmbrnch'>")

	// name
	brncharr = append(brncharr, "<div style='border-top:1px solid rgba(var(--ce3,239,239,239),1);' class='new-branch-div-class-phase'>")
	brncharr = append(brncharr, "<div style='padding-top:10px;' class='new-branch-div-class-phase-title'><span>Add a branch name</span></div>")

	// search for name
	brncharr = append(brncharr, "<div class='new-branch-div-class-phase-srch-name'>")
	brncharr = append(brncharr, "<div id='new-branch-div-class-phase-srch-name01'><button>"+let.Username+" <i data-feather='check' color='black'></i></button></div>")
	brncharr = append(brncharr, "<div id='new-branch-div-class-phase-srch-name02'><span>/</span></div>")
	brncharr = append(brncharr, "<div id='new-branch-div-class-phase-srch-name03' style='position:relative;'><span><input type='text' id='brnch90res671' name='name' placeholder='Branch name' /></span> <div id='fnd-908res'></div></div>")
	brncharr = append(brncharr, "</div>")
	// search for name

	// error name
	brncharr = append(brncharr, "<div id='resname-909sk'></div>")

	// <i style='color:green;height:13px;width:13px;' data-feather='check' color='black'></i>
	// <img style='height:14px;width:14px;' src='/dingtra-web-assets/svg/l.jpg'>

	brncharr = append(brncharr, "</div>")
	// name

	// visibility
	brncharr = append(brncharr, "<div style='display:none;'><span><input type='text' name='visibility' id='visble091res' value='private' /></span></div>")

	// about
	brncharr = append(brncharr, "<div class='new-branch-div-class-phase'>")
	// 
	brncharr = append(brncharr, "<div class='crtxr67090'>")
	brncharr = append(brncharr, AboutBranchLink(r, "about"))
	brncharr = append(brncharr, "</div>")
	// 
	brncharr = append(brncharr, "</div>")
	// about


	// visible
	brncharr = append(brncharr, "<div class='new-branch-div-class-phase'>")
	brncharr = append(brncharr, "<div class='new-branch-div-class-phase-title'><span>Select branch visibility</span></div>")

	// mode of visibility
	brncharr = append(brncharr, "<div class='new-branch-div-class-phase-mode'>")

	brncharr = append(brncharr, `<div id='new-branch-div-class-phase-mode01'>
	<div class="brnchmode0x1"> <input type='checkbox' class="chck_btn0x67" name="private" value='private' checked> <span>Private</span> </div>
	<div class="brnchmode0x2"><span>Only you can add to this branch and make decisions</span></div>
	</div>`)

	brncharr = append(brncharr, `<div id='new-branch-div-class-phase-mode01'>
	<div class="brnchmode0x1"> <input type='checkbox' class="chck_btn0x67" name="public" value='public'> <span>Public</span> </div>
	<div class="brnchmode0x2"><span>Anyone can add to this branch, like adding a new ding, topics and making limited decisions</span></div>
	</div>`)

	brncharr = append(brncharr, "</div>")
	// mode of visibility

	brncharr = append(brncharr, "</div>")
	// visible

	// form
	brncharr = append(brncharr, "</form>")

	// phase submit
	brncharr = append(brncharr, "<div class='new-branch-div-class-phase-submit'>")
	brncharr = append(brncharr, "<div><span><button id='creatbrnchdiv09'>Create branch</button></span></div>")
	brncharr = append(brncharr, "</div>")
	// phase



	brncharr = append(brncharr, "</div></div>")
	// parent


	let.Details = template.HTML(strings.Join(brncharr, ""))
}