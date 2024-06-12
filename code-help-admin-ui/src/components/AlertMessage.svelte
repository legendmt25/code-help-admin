<script lang="ts">
  import { AiOutlineClose } from "svelte-icons-pack/ai";
  import MessageBox from "./MessageBox.svelte";
  import type { MessageBoxType } from "./types";
  import { Icon } from "svelte-icons-pack";

  export let type: MessageBoxType | undefined = undefined;
</script>

<MessageBox {type} style="position:absolute;top:0;" fullwidth>
  <div class="row justify-space-between">
    <slot />
    <Icon size="24" src={AiOutlineClose} />
  </div>
</MessageBox>
