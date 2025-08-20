#include "widget.h"
#include "ui_widget.h"
#include <QString>
#include <vector>
#include <string>
#include <cmath>

using namespace std;

const char argOp[4] = {'+', '-', '*', '/'};
string card[4]{};
double num[4]{};

vector<double> path;
bool used[4]{};
string solution;

// ---------- 算法部分 ----------
void ParseCard() {
    for (int i = 0; i < 4; i++) {
        if (card[i] == "A") num[i] = 1;
        else if (card[i] == "J") num[i] = 11;
        else if (card[i] == "Q") num[i] = 12;
        else if (card[i] == "K") num[i] = 13;
        else if (card[i] == "10") num[i] = 10;
        else if (card[i][0] >= '1' && card[i][0] <= '9') num[i] = card[i][0] - '0';
        else {
            solution = "Invalid Input";
            return;
        }
    }
}

bool operate(double a, double b, char op, double &res) {
    if (op == '+') res = a + b;
    else if (op == '-') res = a - b;
    else if (op == '*') res = a * b;
    else {
        if (fabs(b) < 1e-6) return false;
        res = a / b;
    }
    return true;
}

string tryParentheses(double a, double b, double c, double d, char op1, char op2, char op3) {
    double t1,t2,t3;

    if (operate(a,b,op1,t1) && operate(t1,c,op2,t2) && operate(t2,d,op3,t3) && fabs(t3-24)<1e-6)
        return "((" + to_string((int)a)+op1+to_string((int)b)+")"+op2+to_string((int)c)+")"+op3+to_string((int)d);

    if (operate(b,c,op2,t1) && operate(a,t1,op1,t2) && operate(t2,d,op3,t3) && fabs(t3-24)<1e-6)
        return "(" + to_string((int)a)+op1+"("+to_string((int)b)+op2+to_string((int)c)+"))"+op3+to_string((int)d);

    if (operate(b,c,op2,t1) && operate(t1,d,op3,t2) && operate(a,t2,op1,t3) && fabs(t3-24)<1e-6)
        return to_string((int)a)+op1+"((" + to_string((int)b)+op2+to_string((int)c)+")"+op3+to_string((int)d)+")";

    if (operate(c,d,op3,t1) && operate(b,t1,op2,t2) && operate(a,t2,op1,t3) && fabs(t3-24)<1e-6)
        return to_string((int)a)+op1+"("+to_string((int)b)+op2+"("+to_string((int)c)+op3+to_string((int)d)+"))";

    if (operate(a,b,op1,t1) && operate(c,d,op3,t2) && operate(t1,t2,op2,t3) && fabs(t3-24)<1e-6)
        return "(" + to_string((int)a)+op1+to_string((int)b)+")"+op2+"("+to_string((int)c)+op3+to_string((int)d)+")";

    return "";
}

void calculate() {
    char op[3];
    for(int i = 0;i < 4;i++){
        op[0] = argOp[i];
        for (int j = 0;j < 4;j++) {
            op[1] = argOp[j];
            for (int k = 0;k < 4;k++) {
                op[2] = argOp[k];
                string s = tryParentheses(path[0],path[1],path[2],path[3],op[0],op[1],op[2]);
                if (!s.empty()) {
                    solution = s;
                    return;
                }
            }
        }
    }
}

void dfs() {
    if (path.size()==4) {
        calculate();
        return;
    }
    for (int i=0;i<4 && solution.empty();i++){
        if (!used[i]){
            used[i]=true;
            path.push_back(num[i]);
            dfs();
            path.pop_back();
            used[i]=false;
        }
    }
}
Widget::Widget(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::Widget)
{
    ui->setupUi(this);
}

Widget::~Widget()
{

    delete ui;
}


void Widget::on_Confirm_clicked()
{
    card[0] = ui->Card1->text().toStdString();
    card[1] = ui->Card2->text().toStdString();
    card[2] = ui->Card3->text().toStdString();
    card[3] = ui->Card4->text().toStdString();
    solution.clear();
    path.clear();
    fill(begin(used),end(used),false);

    ParseCard();
    if(solution == "Invalid Input"){
        ui->res->setText("输入错误");
        return;
    }

    dfs();

    if(!solution.empty()){
        ui->res->setText(QString::fromStdString(solution));
    }else{
        ui->res->setText("No Answer");
    }
}

void Widget::on_Exit_clicked()
{
    close();
}

