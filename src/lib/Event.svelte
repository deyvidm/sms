<script lang="ts">
    import { currentUser, pb } from './pocketbase';
  
    let username: string;
    let password: string;
  
    async function login() {
      await pb.collection('users').authWithPassword(username, password);
    }
  
    function newEvent() {
    }
  
</script>


<p>
    <button on:click={newEvent}>New Event</button>
</p>