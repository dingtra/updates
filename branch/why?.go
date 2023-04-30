package branch


import (

)


func (let *BranchStruct) WhyContribute () {
	collect := []string{}

	// parent 
	collect = append(collect, `
	    <div class="why_jss090"><div class="why_jss090-chld">
			<div class="why_jss090_parent">
				<div class="why_jss090_title">
					<span>Reason for contribution</span>
				</div>
	
				<div class="why_jss090_box">
					<span><textarea placeholder="Add reason"></textarea>
				</div>
		   </div>
	    </div></div>
	`)



	collect = append(collect, `
	   </div></div>
	`)
	// parent 
}