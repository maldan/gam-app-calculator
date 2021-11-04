<template>
  <div class="main">
    <!-- <div class="header">
      <div
        v-for="(x, i) in expression.split('')"
        :key="i"
        :class="expression.length - cursor - 1 === i ? 'cursor' : ''"
        v-html="x.replace(' ', '&nbsp;')"
      ></div>
    </div>
    <div class="body">
      <div v-for="x in chars" :key="x" @click="put(x)" class="button clickable">{{ x }}</div>
    </div> -->
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {
    /*document.addEventListener('keydown', (e: KeyboardEvent) => {
      if (e.key === 'Shift') {
        return;
      }
      if (e.key === 'ArrowLeft') this.put('←');
      else if (e.key === 'ArrowRight') this.put('→');
      else if (e.key === 'Delete') this.put('С');
      else if (e.key === 'Backspace') this.put('Del');
      else if (e.key === 'Enter') this.put('=');
      else this.put(e.key);
    });*/
  },
  methods: {
    /*put(c: string) {
      const ITC = (x: number) => x * 2.54;
      const FTC = (x: number) => x * 30.48;
      const LTK = (x: number) => x * 0.453592;

      switch (c) {
        case 'Del':
          const y = this.expression.split('');
          y.splice(this.expression.length - this.cursor - 1, 1);
          this.expression = y.join('');
          break;
        case '→':
          this.cursor -= 1;
          if (this.cursor <= 0) {
            this.cursor = 0;
          }
          break;
        case '←':
          this.cursor += 1;
          if (this.cursor >= this.expression.length) {
            this.cursor = this.expression.length;
          }
          break;
        case 'С':
          this.expression = '';
          this.cursor = 0;
          break;
        case 'Copy':
          console.log(2);
          break;
        case '=':
          try {
            this.expression = eval(this.expression) + '';
          } catch (e) {
            this.expression = e + '';
          }
          break;
        default:
          const x = this.expression.split('');
          x.splice(this.expression.length - this.cursor, 0, c);
          this.expression = x.join('');
      }
    },*/
  },
  data: () => {
    return {
      /*cursor: 0,
      expression: '',
      chars: [
        'С',
        '←',
        '→',
        'Del',
        'Cpy',
        'ITC',
        //
        '1',
        '2',
        '3',
        '+',
        '-',
        'FTC',
        //
        '4',
        '5',
        '6',
        '*',
        '/',
        'LTK',
        //
        '0',
        '.',
        '=',
        '(',
        ')',
        '%',
      ],*/
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  height: 100%;

  .header {
    height: 32px;
    display: flex;
    align-items: center;
    font-size: 32px;
    font-weight: bold;
    padding: 10px;
    justify-content: flex-end;
    position: relative;

    .cursor {
      background: rgba(255, 255, 255, 0.2);
    }
  }

  .body {
    display: grid;
    grid-template-columns: repeat(6, 1fr);
    height: calc(100% - 55px);
    gap: 5px;

    .button {
      flex: 1;
      display: flex;
      align-items: center;
      justify-content: center;
      background: rgba(0, 0, 0, 0.2);
      font-size: 21px;
      font-weight: bold;
      text-transform: uppercase;
    }
  }
}
</style>
