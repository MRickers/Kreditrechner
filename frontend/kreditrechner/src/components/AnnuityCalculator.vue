<template>
    <v-app id="inspire">
      <v-app-bar
        class="px-3"
        color="white"
        flat
        density="compact"
      >
        <v-avatar
          color="grey-darken-1"
          size="32"
        ></v-avatar>
  
        <v-spacer></v-spacer>
  
        <v-tabs
          centered
          color="grey-darken-2"
        >
          <v-tab
            v-for="link in links"
            :key="link"
          >
            {{ link }}
          </v-tab>
        </v-tabs>
        <v-spacer></v-spacer>
  
        <!--v-avatar
          class="hidden-sm-and-down"
          color="grey-darken-1"
          size="32"
        ></v-avatar-->
      </v-app-bar>
  
      <v-main class="bg-grey-lighten-3">
        <v-container>
          <v-row>  
            <v-col
              cols="12"
              sm="8"
            >
              <v-sheet
                class="pa-12"
                min-height="70vh"
                rounded="lg"
              >
              <!-- middle col content  -->
                <v-sheet>
                  <v-row>
                    <v-icon color="blue darken-2" icon="mdi-calculator"></v-icon>
                    <h2 class="text-blue darken-2">Tilgungsrechner</h2>
                  </v-row>
                </v-sheet>
                <v-sheet
                class="py-12"
                >
                    <v-text-field
                    v-model="creditsum"
                    clearable
                    label="Darlehensbetrag"
                    type="text"
                    variant="outlined"
                    :rules="creditsumRules"
                    :counter="10"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihres Darlehens ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-currency-eur"></v-icon>
                      </template>
                    </v-text-field>

                    <v-text-field
                    v-model="interestRate"
                    clearable
                    label="Sollzinssatz"
                    type="text"
                    variant="outlined"
                    :rules="rateRules"
                    :counter="4"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihres Sinssatzes ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-percent"></v-icon>
                      </template>
                    </v-text-field>
                    <v-text-field
                    v-model="repaymentRate"
                    clearable
                    label="Tilgungssatz"
                    type="text"
                    variant="outlined"
                    :rules="rateRules"
                    :counter="4"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihrer anfänglichen geplanten Tilgung ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-percent"></v-icon>
                      </template>
                    </v-text-field>
                    <v-text-field
                    v-model="unscheduledRepaymentRate"
                    clearable
                    label="Sondertilgung"
                    type="text"
                    variant="outlined"
                    :rules="rateRules"
                    :counter="4"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Höhe Ihrer jährlichen geplanten Sondertilgung ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-percent"></v-icon>
                      </template>
                    </v-text-field>
                    <v-text-field
                    v-model="runtime"
                    clearable
                    label="Zinsbindungsdauer"
                    type="text"
                    variant="outlined"
                    :rules="runtimeRules"
                    :counter="2"
                    >
                      <template v-slot:prepend>
                      <v-tooltip location="bottom">
                          <template v-slot:activator="{ props }">
                          <v-icon v-bind="props" icon="mdi-help-circle-outline"></v-icon>
                          </template>

                          Tragen Sie hier die Dauer in Jahren der Zinsbindung ein.
                      </v-tooltip>
                      </template>

                      <template v-slot:append-inner>
                          <v-icon color="blue darken-2" icon="mdi-av-timer"></v-icon>
                      </template>
                    </v-text-field>
                    <v-btn
                  color="success"
                  class="mr-4"
                  @click="calculateAnnuity"
                >
                  Berechnen
                </v-btn>

                <v-table v-if="show_calculation">
                  <thead>
                    <tr>
                      <th class="text-left">
                        Year
                      </th>
                      <th class="text-left">
                        Residual dept [€]
                      </th>
                      <th class="text-left">
                        Interest [€]
                      </th>
                      <th class="text-left">
                        Repayment [€]
                      </th>
                      <th class="text-left">
                        Unsched. Repayment [€]
                      </th>
                      <th class="text-left">
                        Annuity [€]
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="item in calculatedAnnuity.results">
                      <td>{{item.year}}</td>
                      <td>{{item.residualDept}}</td>
                      <td>{{item.interest}}</td>
                      <td>{{item.repayment}}</td>
                      <td>{{item.unscheduledRepayment}}</td>
                      <td>{{item.annuity}}</td>
                    </tr>
                  </tbody>
                </v-table>
                </v-sheet>
                </v-sheet>
            </v-col>
  
            <v-col
              cols="12"
              sm="2"
            >
              <v-sheet
                rounded="lg"
                min-height="268"
              >
                <!--  -->
              </v-sheet>
            </v-col>
          </v-row>
        </v-container>
      </v-main>
    </v-app>
  </template>
  
  <script lang="ts">
  import axios from 'axios';

  class AnnuityResult {
    public year: number;
    public residualDept: number;
    public interest: number;
    public repayment: number;
    public unscheduledRepayment: number;
    public annuity: number;

    public constructor() {
      this.year = 0;
      this.residualDept = 0;
      this.interest = 0;
      this.repayment = 0;
      this.unscheduledRepayment = 0;
      this.annuity = 0;
    }
  }

  class AnnuityResponse {
    public results: AnnuityResult[];

    public constructor() {
      this.results = [];
    }
  }

  class AnnuityRequest {
      public creditsum: number;
      public runtime: number;
      public interestRate: number;
      public repaymentRate: number;
      public unscheduledRepaymentRate: number;

      constructor() {
        this.creditsum = 0;
        this.runtime = 0;
        this.interestRate = 0;
        this.repaymentRate = 0;
        this.unscheduledRepaymentRate = 0;
      }
  }

    export default {
      data: () => ({
        calculatedAnnuity: new AnnuityResponse(),
        show_calculation: false,
        interestRate:'2,5',
        repaymentRate:'2,0',
        unscheduledRepaymentRate:'1,0',
        runtime:'10',
        creditsum: '100000',

        links: [
          'Tilgungsrechner',
          'Bausparrechner'
        ],
        creditsumRules: [
          (v: string | any[]) => !!v || 'Creditsum is required',
          (v: string | any[]) => v.length <= 10 || 'Creditsum must be less than 10 characters',
        ],
        runtimeRules: [
          (v: any) => !!v || 'Runtime is required',
          (v: string | any[]) => v.length <= 2 || 'Runtime must be less than 3 characters',
        ],
        rateRules: [
          (v: any) => !!v || 'Rate is required',
          (v: string | any[]) => v.length <= 4 || 'Rate must be less than 5 characters',
          (v: string) => {
          const pattern_float = /^[0-9]\,[0-9]+$/;
          const pattern_int = /^[0-9]+$/;
          return  pattern_float.test(v) ||
                  pattern_int.test(v) ||
                  'Invalid rate.'
        },
        ],
      }),
      methods: {
        convertToAnnuityRequest(): AnnuityRequest {
          var annuityRequest = new AnnuityRequest();
          annuityRequest.creditsum = parseInt(this.creditsum);
          annuityRequest.interestRate = parseFloat(this.interestRate.replace(',','.'));
          annuityRequest.repaymentRate = parseFloat(this.repaymentRate.replace(',','.'));
          annuityRequest.unscheduledRepaymentRate = parseFloat(this.unscheduledRepaymentRate.replace(',','.'));
          annuityRequest.runtime = parseInt(this.runtime);

          return annuityRequest;
        },
        async calculateAnnuity() {
          this.show_calculation = false;
          var annuityRequest = this.convertToAnnuityRequest();
          var url = window.location.href + "api/v1/calculateAnnuity"
          const {data, status} = await axios.post<AnnuityResponse>(
            url,
            JSON.stringify(annuityRequest),
            {
              headers: {
                ContentType: 'application/json'
              }
            }
          );
          
          if(status != 200) {
            console.log("http status error: ", status);
            return;
          }

          this.calculatedAnnuity = data
          if(this.calculatedAnnuity.results.length <= 0) {
            console.log("invalid annuity: ", this.calculatedAnnuity)
            return;
          }
          this.show_calculation = true;
        }
      }
    }
  </script>