package ui

import x "github.com/plainkit/html"

const tabsJS = `(function(){
  function updateGroup(root, value){
    const triggers = root.querySelectorAll('[data-slot="tabs-trigger"]');
    const contents = root.querySelectorAll('[data-slot="tabs-content"]');
    let activeFound = false;
    triggers.forEach(btn=>{
      const v = btn.getAttribute('data-value')||'';
      const active = v===value;
      btn.setAttribute('aria-selected', active? 'true':'false');
      btn.dataset.state = active? 'active':'inactive';
      if(active) activeFound=true;
      // set tabindex for roving focus
      btn.setAttribute('tabindex', active? '0':'-1');
    });
    contents.forEach(p=>{
      const v = p.getAttribute('data-value')||'';
      const show = v===value;
      p.dataset.state = show? 'active':'inactive';
      if(show) p.removeAttribute('hidden'); else p.setAttribute('hidden','');
    });
    // fallback: if requested value not found, select first
    if(!activeFound && triggers.length){
      const first = triggers[0];
      const v = first.getAttribute('data-value')||'';
      updateGroup(root, v);
    }
  }

  function init(root){
    const triggers = Array.from(root.querySelectorAll('[data-slot="tabs-trigger"]'));
    if(!triggers.length) return;
    // initial value: any marked active, else first
    let current = null;
    for(const btn of triggers){ if(btn.dataset.state==='active'){ current = btn.getAttribute('data-value'); break; } }
    if(!current) current = triggers[0].getAttribute('data-value');
    updateGroup(root, current||'');

    // clicks
    triggers.forEach(btn=>{
      btn.addEventListener('click', e=>{
        e.preventDefault();
        const v = btn.getAttribute('data-value')||'';
        updateGroup(root, v);
        btn.focus();
      });
      // keyboard nav
      btn.addEventListener('keydown', e=>{
        const orient = root.getAttribute('data-orientation')||'horizontal';
        const idx = triggers.indexOf(btn);
        let next = null;
        if((orient==='horizontal' && e.key==='ArrowRight')|| (orient==='vertical' && e.key==='ArrowDown')){
          next = triggers[(idx+1)%triggers.length];
        } else if((orient==='horizontal' && e.key==='ArrowLeft')|| (orient==='vertical' && e.key==='ArrowUp')){
          next = triggers[(idx-1+triggers.length)%triggers.length];
        } else if(e.key==='Home'){ next = triggers[0]; }
        else if(e.key==='End'){ next = triggers[triggers.length-1]; }
        if(next){ e.preventDefault(); const v = next.getAttribute('data-value')||''; updateGroup(root, v); next.focus(); }
      });
    });
  }

  if(document.readyState==='loading'){
    document.addEventListener('DOMContentLoaded',()=>{
      document.querySelectorAll('[data-slot="tabs"]').forEach(init);
    });
  }else{
    document.querySelectorAll('[data-slot="tabs"]').forEach(init);
  }
})();`

// Tabs creates a root tabs container with keyboard navigation and state management
func Tabs(args ...x.DivArg) x.Node {
	base := "flex flex-col gap-2"
	tabsArgs := append([]x.DivArg{
		x.AClass(base),
		x.AData("slot", "tabs"),
	}, args...)

	return x.Div(tabsArgs...).WithAssets("", tabsJS, "tabs")
}

// TabsList creates a container for tab triggers with shadcn/ui styling
func TabsList(args ...x.DivArg) x.Node {
	base := "bg-muted text-muted-foreground inline-flex h-9 w-fit items-center justify-center rounded-lg p-[3px]"
	listArgs := append([]x.DivArg{
		x.AClass(base),
		x.AData("slot", "tabs-list"),
	}, args...)

	return x.Div(listArgs...)
}

// TabsTrigger creates a tab button trigger with shadcn/ui styling. Pass data-value (required) and optional data-state="active"
func TabsTrigger(args ...x.ButtonArg) x.Node {
	base := "cursor-pointer data-[state=active]:bg-background dark:data-[state=active]:text-foreground focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:outline-ring dark:data-[state=active]:border-input dark:data-[state=active]:bg-input/30 text-foreground dark:text-muted-foreground inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap transition-[color,box-shadow] focus-visible:ring-[3px] focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:shadow-sm [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4"

	triggerArgs := append([]x.ButtonArg{
		x.AClass(base),
		x.AData("slot", "tabs-trigger"),
		x.AType("button"),
		x.AAria("selected", "false"),
	}, args...)

	return x.Button(triggerArgs...)
}

// TabsContent creates a tab content panel. Pass data-value to match trigger
func TabsContent(args ...x.DivArg) x.Node {
	base := "flex-1 outline-none"
	contentArgs := append([]x.DivArg{
		x.AClass(base),
		x.AData("slot", "tabs-content"),
		x.AHidden("true"),
	}, args...)

	return x.Div(contentArgs...)
}
