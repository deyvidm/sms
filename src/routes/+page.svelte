<script lang="ts">
    import Header from "$lib/Header.svelte";
    import Sidebar from "$lib/Sidebar.svelte";
    import Contact from "../lib/Contact.svelte";
    import Login from "../lib/Login.svelte";
    import { currentUser } from "../lib/pocketbase";
</script>

<Login />
<Header />
<Sidebar />
