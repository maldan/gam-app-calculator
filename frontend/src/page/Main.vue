<template>
  <div class="main">
    <input
      @keydown.enter="calculate()"
      type="text"
      v-model="expression"
      placeholder="Expression..."
    />
    <div class="history">
      <div
        class="item"
        v-for="x in history"
        :key="x.id"
        v-html="colorize(x.expression + ' = ' + x.result)"
      ></div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {},
  methods: {
    calculate() {
      const now = () => ~~(new Date().getTime() / 1000);
      const inchToCm = (x: number) => x * 2.54;
      const feetToCm = (x: number) => x * 30.48;
      const lbToKg = (x: number) => x * 0.453592;
      const deposit = (amount: number, percentage: number) => {
        const r = percentage / 100;
        amount = amount * r;
        return amount / 12;
      };
      const sin = Math.sin;
      const cos = Math.cos;

      try {
        const expression = this.expression;
        this.expression = eval(this.expression) + '';
        this.history.unshift({ id: Math.random(), expression, result: this.expression });
      } catch (e) {
        this.expression = e + '';
      }
    },
    colorize(expr: string) {
      return expr
        .replace(/(\+|\*|-|\/|=|\(|\))/g, '<span class="operator">$1</span>')
        .replace(/(\d+)/g, '<span class="number">$1</span>');
    },
  },
  data: () => {
    return {
      expression: '',
      history: [] as any[],
    };
  },
});
</script>

<style lang="scss" scoped>
.main {
  height: 100%;
  font-size: 24px;
  box-sizing: border-box;

  input {
    font-size: 24px;
    padding: 20px;
    font-weight: bold;
    background: #1b1b1b;
    color: #999999;
    border: 0;
    outline: none;
    width: 100%;
    box-sizing: border-box;
    margin-bottom: 5px;
  }

  .history {
    height: calc(100% - 80px);
    overflow-y: auto;

    .item {
      padding: 10px;
      background: #1b1b1b;
      color: #ff7a0d;
      font-size: 16px;
      margin-bottom: 5px;
      font-weight: bold;
    }
  }
}
</style>
