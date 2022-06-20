/// A rust implementation made for 2022 Tax records
use clap::Parser;

// Using the Married filing Jointly Standard Deduction
const STANDARD_DED: f64 = 25_900.0;

const TAX_BRACKETS: [(f64, f64); 7] = [
    (647_850.0, 0.37),
    (431_900.0, 0.35),
    (340_100.0, 0.32),
    (178_150.0, 0.24),
    (83_550.0, 0.22),
    (20_550.0, 0.12),
    (0.0, 0.10),
];

#[derive(Debug, Parser)]
#[clap(author, version, about, long_about = None)]
struct Args {
    /// Income to calculate tax burden for
    #[clap(short, long, value_parser)]
    income: f64,
}

fn main() {
    let args = Args::parse();
    if args.income <= 0.0 {
        println!("You can't pay taxes on no income!");
        return;
    }
    let taxable_income = args.income - STANDARD_DED;
    let tax_burden = calc_tax_burden(taxable_income);
    println!("Your taxable income is ${:.2}", taxable_income);
    println!("Your tax burden is ${:.2}", tax_burden);
}

fn calc_tax_burden(mut income: f64) -> f64 {
    let mut total: f64 = 0.0;
    for (bound, rate) in TAX_BRACKETS {
        if income >= bound {
            // We want to calculate the tax for the portion of the income in
            // this Bracket
            total += (income - bound) * rate;
            income = income - (income - bound);
        }
    }
    total
}
