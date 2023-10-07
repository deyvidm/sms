<script>
    import { enhance } from '$app/forms';
</script>

<div class="absolute right-2">
<form use:enhance method="POST" action="?/logout">
    <button class="btn btn-outline-danger">Confirm Loggin Out</button>
</form>
</div>