<template>
    <ul class="List">
        <li>
            <a href="#">
                <span class="material-symbols-outlined">inbox</span>
             Inbox
            </a>
        </li>
        <li>
            <a href="#">
                <span class="material-symbols-outlined">
send
</span>
            Sent Items
            </a>
        </li>
        <li>
            <a href="#">
                <span class="material-symbols-outlined">draft</span>
            Drafts
            </a>
        </li>
    </ul>
</template>

<script>
// Get all the sidebar items
const sidebarItems = document.querySelectorAll('.sidebar-item');

// Add a click event listener to each item
sidebarItems.forEach(item => {
    item.addEventListener('click', () => {
        // 1. Remove the .active class from any element that currently has it
        const currentActive = document.querySelector('.List');
        if (currentActive) {
            currentActive.classList.remove('active');
        }

        // 2. Add the .active class to the item that was just clicked
        item.classList.add('active');
    });
});
</script>

<style>
*{
    padding:0px;
    margin:0px;
}
.List {
  list-style: none;
  padding: 0;
  margin-top: 8px;
}
.List li {
    margin-bottom: 20px;
    padding:5px;

    
}
.List li a{
    color:black;
    text-decoration: none;
    font-size:20px;
    display:flex;
}

.List :hover {
  background: #e3f0ff;
  /* border-radius: 10px; */
  border:none;
  border-radius: 3px;
}
.List :active {
    background: #ccdff6;
   border-radius: 3px;
}

.List :focus {
    background: #ccdff6;
     box-shadow: 0 0 0 4px #ccdff6;
     border-radius: 3px;
  border:none;
}

.material-symbols-outlined{
    margin-right: 6px;
    margin-top:1px;
   font-variation-settings:
    'FILL' 1,   /* 0 = outline, 1 = filled */
    'wght' 300, /* thickness (100–700) */
    'GRAD' 0,   /* grade, affects boldness slightly */
    'opsz' 48;  /* optical size, affects sharpness */
  font-size: 22px; /* size of icon */
}
</style>