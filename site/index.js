$(document).ready(function(){

  var magic8Ball = {};
  magic8Ball.listofanswers = ["kubectl delete namespace", "Tweet @kelseyhightower for help.", "Check github issues.", "It's Docker.", "kubectl cordon ${NODE}", "CrashLoopBackoff, ask question again later.", "Firewalled.", "Throw yourself at the mercy of SIG on Slack.", "Yes.", "Signs point to kubectl get po.", "Reply hazy, try kubectl describe all.", "Ask again later, ImagePullBackOff", "Concentrate and kubectl get events.", "Don't count on it.", "My reply is node is not schedulable.", "My sources say node is tainted.", "Outlook not so healthz.", "Very doubtful. Upgrade cluster to latest version."];

  magic8Ball.getAnswer = function()
  {
    var randomNumber = Math.random();
    var randomAnswer = Math.floor(randomNumber * this.listofanswers.length);
    var answer = this.listofanswers[randomAnswer];

    $("#8ball").effect( "shake" );
    $("#answer").text( answer );
    $("#answer").fadeIn(3000);
    $("#8ball").attr("src", "https://s3.amazonaws.com/media.skillcrush.com/skillcrush/wp-content/uploads/2016/09/answerside.png");

    console.log(question);
    console.log(answer);
  };
  $("#answer").hide();

  var onClick = function()
  {
    $("#answer").hide();
    $("#8ball").attr("src", "https://s3.amazonaws.com/media.skillcrush.com/skillcrush/wp-content/uploads/2016/09/magic8ballQuestion.png");
    magic8Ball.getAnswer();
  };

  $("#questionButton").click( onClick );
});
